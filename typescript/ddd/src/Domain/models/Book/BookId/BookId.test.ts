import { BookId } from './BookId';

describe('BookId', () => {
  test('有効なフォーマットの場合、正しい変換結果を期待', () => {
    expect(new BookId('9784167158057').value).toBe('9784167158057');
    expect(new BookId('4167158051').value)
  });

  test('equals', () => {
    const bookId1 = new BookId('9784167158057');
    const bookId2 = new BookId('9784167158057');
    const bookId3 = new BookId('9785267178086');
    expect(bookId1.equals(bookId2)).toBeTruthy();
    expect(bookId1.equals(bookId3)).toBeFalsy();
  });

  test('toISBN() 13', () => {
    const bookId = new BookId('9784167158057');
    expect(bookId.toISBN()).toBe('ISBN978-4-16-715805-7')
  });

  test('toISBN() 10', () => {
    const bookId = new BookId('4167158057');
    expect(bookId.toISBN()).toBe('ISBN4-16-715805-7')
  });

  test('不正な文字数の場合にエラーを投げる', ()=>{
    expect(()=> new BookId('1'.repeat(101))).toThrow('ISBNの文字数が不正です');
    expect(()=> new BookId('1'.repeat(9))).toThrow('ISBNの文字数が不正です');
  });

  test('不正なフォーマットの場合にエラーを投げる', ()=>{
    expect(()=> new BookId('9991101010101')).toThrow('不正なISBNの形式です');
  });
});
