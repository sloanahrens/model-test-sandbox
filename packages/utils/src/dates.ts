/**
 * Formats a date as ISO string (YYYY-MM-DD).
 */
export function formatDate(date: Date): string {
  return date.toISOString().split('T')[0];
}

/**
 * Parses an ISO date string to a Date object.
 * Returns null if invalid.
 */
export function parseDate(dateStr: string): Date | null {
  const date = new Date(dateStr);
  if (isNaN(date.getTime())) {
    return null;
  }
  return date;
}
