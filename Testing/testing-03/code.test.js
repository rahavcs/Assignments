const { delayedGreeting } = require('./code');

test('delayedGreeting resolves with correct message', async () => {
  const name = 'Alice';
  const delay = 500;
  const result = await delayedGreeting(name, delay);
  expect(result).toBe('Hello, Alice!');
});

test('delayedGreeting respects the delay', () => {
  jest.useFakeTimers();
  const name = 'Bob';
  const delay = 1000;
  const promise = delayedGreeting(name, delay);

  jest.advanceTimersByTime(1000); 

  return promise.then(result => {
    expect(result).toBe('Hello, Bob!');
  });
});
