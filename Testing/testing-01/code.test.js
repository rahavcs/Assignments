// test.js
const { capitalize, reverseString } = require('./code');

describe('String Utilities', () => {

    // Test suite for the capitalize function
    describe('capitalize', () => {
        test('should capitalize the first letter of the word', () => {
            expect(capitalize("hello")).toBe("Hello");
        });

        test('should return an empty string for an empty input', () => {
            expect(capitalize("")).toBe("");
        });

        test('should return the same single-character word', () => {
            expect(capitalize("a")).toBe("A");
        });

        test('should capitalize a word with mixed case letters', () => {
            expect(capitalize("javaScript")).toBe("JavaScript");
        });

        test('should return an empty string for null or undefined input', () => {
            expect(capitalize(null)).toBe("");
            expect(capitalize(undefined)).toBe("");
        });
    });

    // Test suite for the reverseString function
    describe('reverseString', () => {
        test('should reverse a string correctly', () => {
            expect(reverseString("hello")).toBe("olleh");
        });

        test('should return an empty string when the input is empty', () => {
            expect(reverseString("")).toBe("");
        });

        test('should return the same string for a single character', () => {
            expect(reverseString("a")).toBe("a");
        });

        test('should correctly reverse a palindrome', () => {
            expect(reverseString("madam")).toBe("madam");
        });

        test('should correctly reverse a string with special characters', () => {
            expect(reverseString("hello!")).toBe("!olleh");
        });

        test('should reverse a string with spaces', () => {
            expect(reverseString("hello world")).toBe("dlrow olleh");
        });

        test('should handle long strings correctly', () => {
            const longString = "a".repeat(1000); // Create a string of 1000 'a's
            expect(reverseString(longString)).toBe(longString.split("").reverse().join(""));
        });
    });

});
