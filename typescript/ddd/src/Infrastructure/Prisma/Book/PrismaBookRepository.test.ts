import { PrismaClient } from "@prisma/client";
import { PrismaBookRepository } from "./PrismaBookRepository";
import { BookId } from "Domain/models/Book/BookId/BookId";
import { Title } from "Domain/models/Book/Title/Title";
import { Price } from "Domain/models/Book/Price/Price";
import { Book } from "Domain/models/Book/Book";
import { bookTestDataCreater } from "./bookTestDataCreator";
import { Stock } from "Domain/models/Book/Stock/Stock";
import { QuantityAvailable } from "Domain/models/Book/Stock/QuantityAvailable/QuantityAvailable";
import { Status, StatusEnum } from "Domain/models/Book/Stock/Status/Status";
import { PrismaClientManager } from "../PrismaClientManager";

const prisma = new PrismaClient();

describe('PrismaBookRepository', () => {
  beforeEach(async () => {
    await prisma.$transaction([prisma.book.deleteMany()]);
    await prisma.$disconnect();
  });

  const clientManager = new PrismaClientManager();
  const repository = new PrismaBookRepository(clientManager);

  test('saveした集約がfindで取得できる', async () => {
    const bookId = new BookId('9781111111111');
    const title = new Title('吾輩は猫である');
    const price = new Price({
      amount: 770,
      currency: 'JPY',
    });
    const book = Book.create(bookId, title, price);
    await repository.save(book);

    const createEntity = await repository.find(bookId);
    expect(createEntity?.bookId.equals(bookId)).toBeTruthy();
    expect(createEntity?.title.equals(title)).toBeTruthy();
    expect(createEntity?.price.equals(price)).toBeTruthy();
    expect(createEntity?.stockId.equals(book.stockId)).toBeTruthy();
    expect(createEntity?.quantityAvailable.equals(book.quantityAvailable)).toBeTruthy();
    expect(createEntity?.status.equals(book.status)).toBeTruthy();
  });

  test('updateできる', async () => {
    const createdEntity = await bookTestDataCreater(repository)({});

    const stock = Stock.reconstruct(
      createdEntity.stockId,
      new QuantityAvailable(100),
      new Status(StatusEnum.InStock)
    );

    const book = Book.reconstruct(
      createdEntity.bookId,
      new Title('吾輩は猫である(改訂版)'),
      new Price({ amount: 880, currency: 'JPY' }),
      stock
    );

    await repository.update(book);
    const updatedEntity = await repository.find(createdEntity.bookId);
    expect(updatedEntity?.bookId.equals(book.bookId)).toBeTruthy();
    expect(updatedEntity?.title.equals(book.title)).toBeTruthy();
    expect(updatedEntity?.price.equals(book.price)).toBeTruthy();
    expect(updatedEntity?.stockId.equals(book.stockId)).toBeTruthy();
    expect(updatedEntity?.quantityAvailable.equals(book.quantityAvailable)).toBeTruthy();
    expect(updatedEntity?.status.equals(book.status)).toBeTruthy();
  });

  test('deleteできる', async () => {
    const createdEntity = await bookTestDataCreater(repository)({});

    const readEntity = await repository.find(createdEntity.bookId);
    expect(readEntity).not.toBeNull();

    await repository.delete(createdEntity.bookId);
    const deleteEntity = await repository.find(createdEntity.bookId);
    expect(deleteEntity).toBeNull();
  });
});
