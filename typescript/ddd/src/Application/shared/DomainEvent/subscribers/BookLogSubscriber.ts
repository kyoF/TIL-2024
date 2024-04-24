import { BOOK_EVENT_NAME, BookDomainEventBody } from "Domain/shared/DomainEvent/Book/BookDomainEventFactory";
import { IDomainEvnetSubscriber } from "Domain/shared/DomainEvent/IDomainEventSubscriber";
import { inject, injectable } from "tsyringe";

@injectable()
export class BookLogSubscriber {
  constructor(
    @inject('IDomainEventSubscriber')
    private subscriber: IDomainEvnetSubscriber
  ) {
    this.subscriber.subscribe<BookDomainEventBody>(
      BOOK_EVENT_NAME.CREATED,
      (event) => {
        console.log(event);
      }
    );
  }
}
