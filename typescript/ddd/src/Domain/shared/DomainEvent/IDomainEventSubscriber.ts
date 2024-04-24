import { DomainEvent } from "./DomainEvent";

export interface IDomainEvnetSubscriber {
  subscribe<T extends Record<string, unknown>>(
    eventName: string,
    callback: (event: DomainEvent<T>) => void
  ): void;
}
