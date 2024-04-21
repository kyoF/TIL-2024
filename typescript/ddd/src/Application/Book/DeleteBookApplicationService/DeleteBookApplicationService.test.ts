import { MockTransactionManager } from "Application/shared/MockTransactionManager";
import { InMemoryBookRepository } from "Infrastructure/InMemory/Book/InMemoryBookRepository";
import { DeleteBookApplicationService, DeleteBookCommand } from "./DeleteBookApplicationService";
import { bookTestDataCreater } from "Infrastructure/Prisma/Book/bookTestDataCreator";
import { BookId } from "Domain/models/Book/BookId/BookId";

describe('DeleteBookApplicationService', () => {
  test('書籍を削除することができる', async () => {
    const repository = new InMemoryBookRepository();
    const mockTransactionManager = new MockTransactionManager();
    const deleteBookApplicationService = new DeleteBookApplicationService(
      repository,
      mockTransactionManager
    );

    const bookId = '9781111111111';
    await bookTestDataCreater(repository)({
      bookId,
    });

    const command: Required<DeleteBookCommand> = { bookId };
    await deleteBookApplicationService.execute(command);

    const deleteBook = await repository.find(new BookId(bookId));
    expect(deleteBook).toBe(null);
  });
});
