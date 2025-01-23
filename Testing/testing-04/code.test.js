const { sendNotification } = require('./code');

// Test 1: Success
test('Notification is sent successfully', () => {
  const mockService = { send: () => true };  
  const result = sendNotification(mockService, 'Hello');
  expect(result).toBe('Notification Sent');
});

// Test 2: Failure
test('Notification failed to send', () => {
  const mockService = { send: () => false };
  const result = sendNotification(mockService, 'Hello');
  expect(result).toBe('Failed to Send');
});
