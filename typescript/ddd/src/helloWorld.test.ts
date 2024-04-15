import { helloWorld } from './helloWorld';

test('say-hello-world', () => {
  expect(helloWorld('world')).toBe('Hello world!');
});
