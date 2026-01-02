import type { Logger } from './logger.js';

export interface RetryOptions {
  maxAttempts: number;
  delayMs: number;
  backoff?: boolean;
  logger?: Logger;
}

/**
 * Sleep for a specified number of milliseconds.
 */
export function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * Retry a function with exponential backoff.
 */
export async function retry<T>(
  fn: () => Promise<T>,
  options: RetryOptions
): Promise<T> {
  const { maxAttempts, delayMs, backoff = true, logger } = options;

  let lastError: Error | undefined;

  for (let attempt = 1; attempt <= maxAttempts; attempt++) {
    try {
      return await fn();
    } catch (error) {
      lastError = error instanceof Error ? error : new Error(String(error));
      logger?.warn(`Attempt ${attempt}/${maxAttempts} failed: ${lastError.message}`);

      if (attempt === maxAttempts) {
        break;
      }

      const waitTime = backoff ? delayMs * Math.pow(2, attempt - 1) : delayMs;
      logger?.debug(`Waiting ${waitTime}ms before retry`);
      await sleep(waitTime);
    }
  }

  logger?.error(`All ${maxAttempts} attempts failed`);
  throw lastError;
}
