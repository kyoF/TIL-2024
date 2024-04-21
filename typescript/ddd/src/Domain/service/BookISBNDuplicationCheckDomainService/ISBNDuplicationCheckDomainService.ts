import { BookId } from 'Domain/models/Book/BookId/BookId';

export class ISBNDuplicationCheckDomainService {
  async execute(_ /* bookId */: BookId): Promise<boolean> {
    const isDuplicateISBN = false;
    return isDuplicateISBN;
  }
}
