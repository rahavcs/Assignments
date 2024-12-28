function parseJSONData(jsonData) {
    try {
        const products = JSON.parse(jsonData); 
        return products;
    } catch (error) {
        console.error("Error parsing JSON:", error);
    }
}


function addProduct(products, newProduct) {
    products.push(newProduct);
    return products;
}


function updateProductPrice(products, productId, newPrice) {
    const product = products.find(product => product.id === productId);
    
    if (product) {
        product.price = newPrice;
        return products;
    } else {
        return "Product not found";
    }
}


function filterAvailableProducts(products) {
    return products.filter(product => product.available === true);
}


function filterProductsByCategory(products, categoryName) {
    return products.filter(product => product.category === categoryName);
}


const jsonData = `[
    {"id": 1, "name": "Laptop", "category": "Electronics", "price": 999.99, "available": true},
    {"id": 2, "name": "Smartphone", "category": "Electronics", "price": 599.99, "available": false},
    {"id": 3, "name": "Coffee Maker", "category": "Home Appliances", "price": 79.99, "available": true}
]`;

let products = parseJSONData(jsonData);

// Add a new product
const newProduct = {"id": 4, "name": "Washing Machine", "category": "Home Appliances", "price": 499.99, "available": true};
products = addProduct(products, newProduct);

// Update the price of a product
products = updateProductPrice(products, 2, 549.99);

// Filter available products
const availableProducts = filterAvailableProducts(products);

// Filter products by category
const electronics = filterProductsByCategory(products, "Electronics");

// results
console.log("Available Products:", availableProducts);
console.log("Electronics:", electronics);
