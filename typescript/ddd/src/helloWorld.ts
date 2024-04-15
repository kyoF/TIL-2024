export const helloWorld = (name: string): string => {
  const res = `Hello ${name}!`;
  console.log(res);
  return res;
}

helloWorld('world');
