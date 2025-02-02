const delayedGreeting = require('../src/delayedGreeting');

jest.useFakeTimers();

describe('delayedGreeting', () => {
  beforeEach(() => {
    jest.clearAllTimers();
  });

  test('should resolve with the correct greeting message', async () => {
    const name = 'Sakshi';
    const delay = 1000;

    const promise = delayedGreeting(name, delay);

    // Fast-forward the timer
    jest.advanceTimersByTime(delay);
    const result = await promise;

    expect(result).toBe(`Hello, ${name}!`);
  });

  test('should respect the delay time', () => {
    const name = 'Sakshi';
    const delay = 1000;

    // Spy on setTimeout
    const setTimeoutSpy = jest.spyOn(global, 'setTimeout');

    delayedGreeting(name, delay);

    // Ensure setTimeout was called with the correct delay
    expect(setTimeoutSpy).toHaveBeenCalledWith(expect.any(Function), delay);

    // Restore the spy after the test
    setTimeoutSpy.mockRestore();
  });
});
