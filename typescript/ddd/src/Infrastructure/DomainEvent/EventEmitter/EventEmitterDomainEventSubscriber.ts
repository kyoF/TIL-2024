import { DomainEvent } from "Domain/shared/DomainEvent/DomainEvent";
import { IDomainEvnetSubscriber } from "Domain/shared/DomainEvent/IDomainEventSubscriber";
import { container } from "tsyringe";
import EventEmitterClient from "./EventEmitterClient";

export class EventEmitterDomainEventSubscriber implements IDomainEvnetSubscriber {
  subscribe<T extends Record<string, unknown>>(
    eventName: string,
    callback: (event: DomainEvent<T>) => void
  ) {
    container.resolve(EventEmitterClient)
      .eventEmitter.once(eventName, callback);
  }
}
