import { InMemoryBookRepository } from "Infrastructure/InMemory/Book/InMemoryBookRepository";
import { GetBookApplicationService } from "./GetBookApplicationService";
import { bookTestDataCreater } from "Infrastructure/Prisma/Book/bookTestDataCreator";
import { BookDTO } from "../BookDTO";

describe('GetBookApplicationService', () => {
  test('指定されたIDの書籍が存在する場合、DTOに詰め替えられ、取得できる', async () => {
    const repository = new InMemoryBookRepository();
    const getBookApplicationService = new GetBookApplicationService(repository);

    const createdBook = await bookTestDataCreater(repository)({});

    const data = await getBookApplicationService.execute(
      createdBook.bookId.value
    );

    expect(data).toEqual(new BookDTO(createdBook));
  });

  test('指定されたIDの書籍が存在しない場合、nullが取得できる', async () => {
    const repository = new InMemoryBookRepository();
    const getBookApplicationService = new GetBookApplicationService(repository);

    const data = await getBookApplicationService.execute('9781111111111');

    expect(data).toBeNull();
  });
});
