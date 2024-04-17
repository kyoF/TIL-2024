import { Status, StatusEnum } from './Status';

describe('Status', () => {
  test('有効なステータスでインスタンスが生成されること', () => {
    expect(new Status(StatusEnum.InStock).value).toBe(StatusEnum.InStock);
    expect(new Status(StatusEnum.LowStock).value).toBe(StatusEnum.LowStock);
    expect(new Status(StatusEnum.OutOfStock).value).toBe(StatusEnum.OutOfStock);
  });

  test('無効なステータスでエラーが投げられること', () => {
    const invalidStatus = 'invalid' as StatusEnum;
    expect(() => new Status(invalidStatus)).toThrow('無効なステータスです');
  });
});

describe('toLabel()', () => {
  test('ステータスInStockが「在庫あり」に変換されること', () => {
    const status = new Status(StatusEnum.InStock);
    expect(status.toLabel()).toBe('在庫あり');
  });

  test('ステータスOutOfStockが「在庫切れ」に変換されること', () => {
    const status = new Status(StatusEnum.OutOfStock);
    expect(status.toLabel()).toBe('在庫切れ');
  });

  test('ステータスLowStockが「残りわずか」に変換されること', () => {
    const status = new Status(StatusEnum.LowStock);
    expect(status.toLabel()).toBe('残りわずか');
  });
});
