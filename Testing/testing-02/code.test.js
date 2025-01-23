const { getElement } = require('./code'); 

describe('getElement function', () => {
    test('returns the element at a valid index', () => {
        expect(getElement([1, 2, 3], 1)).toBe(2);  
    });
    test('throws error for a negative index', () => {
        expect(() => getElement([1, 2, 3], -1)).toThrow("Index out of bounds");
    });
    test('throws error for an out of range index', () => {
        expect(() => getElement([1, 2, 3], 5)).toThrow("Index out of bounds");
    });
});
