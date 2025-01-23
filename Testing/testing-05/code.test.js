const { toggleVisibility } = require('./code');

test('should hide element when it is visible', () => {
  const element = { style: { display: 'block' } };
  
  toggleVisibility(element);
  
  expect(element.style.display).toBe('none');
});

test('should show element when it is hidden', () => {
  const element = { style: { display: 'none' } };
  
  toggleVisibility(element);
  
  expect(element.style.display).toBe('block');
});
