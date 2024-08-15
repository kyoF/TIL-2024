import { Ref } from 'vue';

export const useTitle = () => {
  const title = useState<string>('title', () => 'Hello World!');
  const changeTitle = (title: Ref<string>) => (value: string) => {
    title.value = value;
  };
  return {
    title: readonly(title),
    changeTitle: changeTitle(title),
  }
};

export const useCounter = () => {
  const num = useState<number>('num', () => 0);
  const countUp = (num: Ref<number>) => () => {
    num.value += 1;
  };
  const countDown = (num: Ref<number>) => () => {
    num.value -= 1;
  };
  const countReset = (num: Ref<number>) => () => {
    num.value = 0;
  };
  return {
    num: readonly(num),
    countUp: countUp(num),
    countDown: countDown(num),
    countReset: countReset(num),
  }
}
