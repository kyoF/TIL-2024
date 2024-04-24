import { DomainEvent } from 'Domain/shared/DomainEvent/DomainEvent';
import { IDomainEvnetSubscriber } from 'Domain/shared/DomainEvent/IDomainEventSubscriber';

export class MockDomainEventSubscriber implements IDomainEvnetSubscriber {
  subscribe<T extends Record<string, unknown>>(
    eventName: string,
    callback: (event: DomainEvent<T>) => void
  ) {
    callback;
    eventName;
  }
}
