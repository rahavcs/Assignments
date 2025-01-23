// code.js

// Capitalize the first letter of the string
function capitalize(word) {
    if (!word) return "";
    return word[0].toUpperCase() + word.slice(1);
}

// Reverse a string
function reverseString(str) {
    return str.split("").reverse().join("");
}

// Export the functions
module.exports = { capitalize, reverseString };
