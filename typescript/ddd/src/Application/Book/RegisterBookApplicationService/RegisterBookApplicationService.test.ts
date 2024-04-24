import { MockTransactionManager } from "Application/shared/MockTransactionManager";
import { InMemoryBookRepository } from "Infrastructure/InMemory/Book/InMemoryBookRepository";
import { RegisterBookApplicatoinService, RegisterBookCommand } from "./RegisterBookApplicationService";
import { BookId } from "Domain/models/Book/BookId/BookId";
import { bookTestDataCreater } from "Infrastructure/Prisma/Book/bookTestDataCreator";
import { MockDomainEventPublisher } from "Infrastructure/DomainEvent/Mock/MockDomainEventPublisher";

describe('RegisterBookApplicationService', () => {
  test('重複書籍が存在しない場合、書籍が正常に作成できる', async () => {
    const repository = new InMemoryBookRepository();
    const mockTransactionManager = new MockTransactionManager();
    const mockDomainEventPublisher = new MockDomainEventPublisher();
    const registerBookApplicationService = new RegisterBookApplicatoinService(
      repository,
      mockTransactionManager,
      mockDomainEventPublisher,
    );

    const command: Required<RegisterBookCommand> = {
      isbn: '9781111111111',
      title: '吾輩は猫である',
      priceAmount: 770,
    };

    await registerBookApplicationService.execute(command);

    const createdBook = await repository.find(new BookId(command.isbn));
    expect(createdBook).not.toBeNull();
  });

  test('重複書籍が存在する場合、エラーを投げる', async () => {
    const repository = new InMemoryBookRepository();
    const mockTransactionManager = new MockTransactionManager();
    const mockDomainEventPublisher = new MockDomainEventPublisher();
    const registerBookApplicationService = new RegisterBookApplicatoinService(
      repository,
      mockTransactionManager,
      mockDomainEventPublisher,
    );

    const bookId = '9781111111111'
    await bookTestDataCreater(repository)({
      bookId: bookId,
    });

    const command: Required<RegisterBookCommand> = {
      isbn: bookId,
      title: '吾輩は猫である',
      priceAmount: 770,
    };

    await expect(
      registerBookApplicationService.execute(command)
    ).rejects.toThrow();
  });
});
