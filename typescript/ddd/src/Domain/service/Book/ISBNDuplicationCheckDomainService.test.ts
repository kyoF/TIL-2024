import { InMemoryBookRepository } from "Infrastructure/InMemory/Book/InMemoryBookRepository";
import { ISBNDuplicationCheckDomainService } from "./ISBNDuplicationCheckDomainService";
import { BookId } from "Domain/models/Book/BookId/BookId";
import { Title } from "Domain/models/Book/Title/Title";
import { Price } from "Domain/models/Book/Price/Price";
import { Book } from "Domain/models/Book/Book";

describe('ISBNDuplicationCheckDomainService', () => {
  let isbnDuplicationCheckDomainService: ISBNDuplicationCheckDomainService;
  let inMemoryBookRepository: InMemoryBookRepository;

  beforeEach(() => {
    inMemoryBookRepository = new InMemoryBookRepository();
    isbnDuplicationCheckDomainService = new ISBNDuplicationCheckDomainService(
      inMemoryBookRepository
    );
  });

  test('重複がない場合、falseを返す', async () => {
    const isbn = new BookId('9781111111111');
    const result = await isbnDuplicationCheckDomainService.execute(isbn);
    expect(result).toBeFalsy();
  });

  test('重複がある場合、trueを返す', async () => {
    const isbn = new BookId('9781111111111');
    const title = new Title('吾輩は猫である');
    const price = new Price({
      amount: 770,
      currency: 'JPY'
    });
    const book = Book.create(isbn, title, price);

    await inMemoryBookRepository.save(book);

    const result = await isbnDuplicationCheckDomainService.execute(isbn);
    expect(result).toBeTruthy();
  });

  test('異なるISBNで重複がない場合、falseを返す', async () => {
    const existingISBN = new BookId('9781111111111');
    const newIsbn = new BookId('9781111111112');
    const title = new Title('吾輩は猫である');
    const price = new Price({
      amount: 770,
      currency: 'JPY'
    });
    const book = Book.create(existingISBN, title, price);

    await inMemoryBookRepository.save(book);

    const result = await isbnDuplicationCheckDomainService.execute(newIsbn);
    expect(result).toBeFalsy();
  });
});
