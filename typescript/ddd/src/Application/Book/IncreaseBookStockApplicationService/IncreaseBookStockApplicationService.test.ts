import { MockTransactionManager } from "Application/shared/MockTransactionManager";
import { InMemoryBookRepository } from "Infrastructure/InMemory/Book/InMemoryBookRepository";
import { IncreaseBookStockApplicationService, IncreaseBookStockCommand } from "./IncreaseBookStockApplicationService";
import { bookTestDataCreater } from "Infrastructure/Prisma/Book/bookTestDataCreator";
import { BookId } from "Domain/models/Book/BookId/BookId";

describe('IncreaseBookStockApplicationService', () => {
  test('書籍の在庫を増加することができる', async () => {
    const repository = new InMemoryBookRepository();
    const mockTransactionManager = new MockTransactionManager();
    const increaseBookStockApplicationService = new IncreaseBookStockApplicationService(
      repository, mockTransactionManager
    );

    const bookId = '9781111111111';
    await bookTestDataCreater(repository)({
      bookId,
      quantityAvailable: 0,
    });

    const incrementAmount = 100;
    const command: Required<IncreaseBookStockCommand> = {
      bookId,
      incrementAmount,
    };
    await increaseBookStockApplicationService.execute(command);

    const updatedBook = await repository.find(new BookId(bookId));
    expect(updatedBook?.quantityAvailable.value).toBe(incrementAmount);
  });

  test('書籍が存在しない場合、エラーを投げる', async () => {
    const repository = new InMemoryBookRepository();
    const mockTransactionManager = new MockTransactionManager();
    const increaseBookStockApplicationService = new IncreaseBookStockApplicationService(
      repository, mockTransactionManager
    );

    const bookId = '9781111111111';
    const incrementAmount = 100;
    const command: Required<IncreaseBookStockCommand> = {
      bookId,
      incrementAmount,
    };

    await expect(
      increaseBookStockApplicationService.execute(command)
    ).rejects.toThrow();
  });
});
