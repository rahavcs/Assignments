const { fetchAndDisplayUser } = require('./code');

// Test if it works when API is successful
test('displays user name if API works', async () => {
  const mockApiService = {
    getUser: jest.fn().mockResolvedValue({ name: 'John Doe' }), 
  };

  const element = { textContent: '' };
  const userId = 1;

  await fetchAndDisplayUser(mockApiService, userId, element);

  expect(element.textContent).toBe('Hello, John Doe'); 
});

// Test if it works when API fails
test('displays error message if API fails', async () => {
  const mockApiService = {
    getUser: jest.fn().mockRejectedValue(new Error('User not found')), 
  };

  const element = { textContent: '' };
  const userId = 1;

  await fetchAndDisplayUser(mockApiService, userId, element);

  expect(element.textContent).toBe('User not found'); 
});

// Test if it works when API returns invalid user data
test('displays error message if user data is invalid', async () => {
  const mockApiService = {
    getUser: jest.fn().mockResolvedValue({}), 
  };

  const element = { textContent: '' };
  const userId = 1;

  await fetchAndDisplayUser(mockApiService, userId, element);

  expect(element.textContent).toBe('Invalid user data'); 
});
