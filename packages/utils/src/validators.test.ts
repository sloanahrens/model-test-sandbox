import { describe, it, expect } from 'vitest';
import { validateEmail, validateUrl } from './validators.js';

describe('validateEmail', () => {
  it('accepts valid emails', () => {
    expect(validateEmail('user@example.com')).toBe(true);
    expect(validateEmail('user.name@domain.org')).toBe(true);
  });

  it('rejects invalid emails', () => {
    expect(validateEmail('invalid')).toBe(false);
    expect(validateEmail('@nodomain.com')).toBe(false);
    expect(validateEmail('spaces in@email.com')).toBe(false);
  });
});

describe('validateUrl', () => {
  it('accepts valid URLs', () => {
    expect(validateUrl('https://example.com')).toBe(true);
    expect(validateUrl('http://localhost:3000')).toBe(true);
  });

  it('rejects invalid URLs', () => {
    expect(validateUrl('not-a-url')).toBe(false);
    expect(validateUrl('')).toBe(false);
  });
});
