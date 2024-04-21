import { Book } from "Domain/models/Book/Book";
import { StatusLable } from "Domain/models/Book/Stock/Status/Status";

export class BookDTO {
  public readonly bookId: string;
  public readonly title: string;
  public readonly price: number;
  public readonly stockId: string;
  public readonly quantityAvailable: number;
  public readonly status: StatusLable;

  constructor(book: Book) {
    this.bookId = book.bookId.value;
    this.title = book.title.value;
    this.price = book.price.amount
    this.stockId = book.stockId.value
    this.quantityAvailable = book.quantityAvailable.value;
    this.status = book.status.toLabel();
  }
}
