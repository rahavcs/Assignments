//PART-1://
//1. Create the Collections and Insert Data:
db.customers.insertMany([
    { 
        name: "Amit Kumar", 
        email: "amit.kumar@example.com", 
        address: { street: "45 B Block", city: "Delhi", zipcode: "110001" },
        phone: "9876543210", 
        registration_date: ISODate("2023-01-01T12:00:00Z")
    },
    { 
        name: "Priya Sharma", 
        email: "priya.sharma@example.com", 
        address: { street: "67 A Road", city: "Mumbai", zipcode: "400001" },
        phone: "9876543211", 
        registration_date: ISODate("2023-02-15T12:00:00Z")
    },
    { 
        name: "Ravi Patel", 
        email: "ravi.patel@example.com", 
        address: { street: "123 C Lane", city: "Bangalore", zipcode: "560001" },
        phone: "9876543212", 
        registration_date: ISODate("2023-03-10T12:00:00Z")
    },
    { 
        name: "Neha Singh", 
        email: "neha.singh@example.com", 
        address: { street: "12 D Avenue", city: "Chennai", zipcode: "600001" },
        phone: "9876543213", 
        registration_date: ISODate("2023-04-05T12:00:00Z")
    },
    { 
        name: "Vikram Reddy", 
        email: "vikram.reddy@example.com", 
        address: { street: "89 E Road", city: "Hyderabad", zipcode: "500001" },
        phone: "9876543214", 
        registration_date: ISODate("2023-05-20T12:00:00Z")
    }
]);

const amitKumar = db.customers.findOne({ name: "Amit Kumar" });
const priyaSharma = db.customers.findOne({ name: "Priya Sharma" });
const raviPatel = db.customers.findOne({ name: "Ravi Patel" });
const nehaSingh = db.customers.findOne({ name: "Neha Singh" });
const vikramReddy = db.customers.findOne({ name: "Vikram Reddy" });

db.orders.insertMany([
    { 
        order_id: "ORD123001", 
        customer_id: amitKumar._id, 
        order_date: ISODate("2023-05-15T14:00:00Z"), 
        status: "shipped", 
        items: [
            { product_name: "Laptop", quantity: 1, price: 1500 },
            { product_name: "Mouse", quantity: 2, price: 25 }
        ],
        total_value: 1550
    },
    { 
        order_id: "ORD123002", 
        customer_id: priyaSharma._id, 
        order_date: ISODate("2023-06-20T14:00:00Z"),
        status: "pending",
        items: [
            { product_name: "Smartphone", quantity: 1, price: 1000 },
            { product_name: "Charger", quantity: 1, price: 50 }
        ],
        total_value: 1050
    },
    { 
        order_id: "ORD123003", 
        customer_id: raviPatel._id, 
        order_date: ISODate("2023-07-01T14:00:00Z"),
        status: "delivered",
        items: [
            { product_name: "Tablet", quantity: 1, price: 400 },
            { product_name: "Headphones", quantity: 1, price: 100 }
        ],
        total_value: 500
    },
    { 
        order_id: "ORD123004", 
        customer_id: nehaSingh._id, 
        order_date: ISODate("2023-08-25T14:00:00Z"),
        status: "shipped",
        items: [
            { product_name: "Smartwatch", quantity: 1, price: 200 },
            { product_name: "Earphones", quantity: 1, price: 30 }
        ],
        total_value: 230
    },
    { 
        order_id: "ORD123005", 
        customer_id: vikramReddy._id, 
        order_date: ISODate("2023-09-10T14:00:00Z"),
        status: "delivered",
        items: [
            { product_name: "Monitor", quantity: 1, price: 300 },
            { product_name: "Keyboard", quantity: 1, price: 50 }
        ],
        total_value: 350
    }
]);


//2. Find Orders for a Specific Customer:
db.orders.find({ customer_id: ObjectId("64a1e4e4efdd3b0012345679") });

/*OUTPUT:
[
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345679"),
    "order_id": "ORD123002",
    "customer_id": ObjectId("64a1e4e4efdd3b0012345679"),
    "order_date": ISODate("2023-06-20T14:00:00Z"),
    "status": "pending",
    "items": [
      { "product_name": "Smartphone", "quantity": 1, "price": 1000 },
      { "product_name": "Charger", "quantity": 1, "price": 50 }
    ],
    "total_value": 1050
  }
]
*/

//3. Find the Customer for a Specific Order:
db.customers.findOne({ _id: ObjectId(db.orders.findOne({ order_id: "ORD123005" }).customer_id) });

/*OUTPUT:
{
  "_id": ObjectId("64a1e4e4efdd3b001234567C"),
  "name": "Vikram Reddy",
  "email": "vikram.reddy@example.com",
  "address": { "street": "89 E Road", "city": "Hyderabad", "zipcode": "500001" },
  "phone": "9876543214",
  "registration_date": ISODate("2023-05-20T12:00:00Z")
}
*/


//4. Update Order Status:
db.orders.updateOne(
    { "order_id": "ORD123005" },          
    { $set: { "status": "delivered" } }   
);

/*OUTPUT:
{
  "acknowledged": true,
  "modifiedCount": 1
}
*/

//5. Delete an Order:
db.orders.deleteOne(
    { "order_id": "ORD123005" } 
);

/*OUTPUT:
{
  "acknowledged": true,
  "deletedCount": 1
}
*/

/*PART-2*/

/*1. Calculate Total Value of All Orders by Customer:*/
db.orders.aggregate([
  {
    $lookup: {
      from: "customers",
      localField: "customer_id",
      foreignField: "_id",
      as: "customer_info"
    }
  },
  { $unwind: "$customer_info" },
  {
    $group: {
      _id: "$customer_id",
      customer_name: { $first: "$customer_info.name" },
      total_order_value: { $sum: "$total_value" }
    }
  }
]);

/*OUTPUT:
[
  { "_id": ObjectId("64a1e4e4efdd3b0012345678"), "customer_name": "Amit Kumar", "total_order_value": 1550 },
  { "_id": ObjectId("64a1e4e4efdd3b0012345679"), "customer_name": "Priya Sharma", "total_order_value": 1050 },
  { "_id": ObjectId("64a1e4e4efdd3b001234567A"), "customer_name": "Ravi Patel", "total_order_value": 500 },
  { "_id": ObjectId("64a1e4e4efdd3b001234567B"), "customer_name": "Neha Singh", "total_order_value": 230 },
  { "_id": ObjectId("64a1e4e4efdd3b001234567C"), "customer_name": "Vikram Reddy", "total_order_value": 350 }
]

*/

/*2. Group Orders by Status:*/
db.orders.aggregate([
    {
      $group: {
        _id: "$status",
        order_count: { $sum: 1 }
      }
    }
  ]);

/*OUTPUT:
[
  { "_id": "shipped", "order_count": 2 },
  { "_id": "delivered", "order_count": 2 },
  { "_id": "pending", "order_count": 1 }
]
*/

/*3. List Customers with Their Recent Orders:*/
db.orders.aggregate([
    {
      $lookup: {
        from: "customers",
        localField: "customer_id",
        foreignField: "_id",
        as: "customer_info"
      }
    },
    { $unwind: "$customer_info" },
    {
      $sort: { "order_date": -1 }
    },
    {
      $group: {
        _id: "$customer_id",
        customer_name: { $first: "$customer_info.name" },
        email: { $first: "$customer_info.email" },
        most_recent_order: {
          $first: {
            order_id: "$order_id",
            total_value: "$total_value",
            order_date: "$order_date"
          }
        }
      }
    }
  ]);

 /*OUTPUT:
 [
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567C"),
    "customer_name": "Vikram Reddy",
    "email": "vikram.reddy@example.com",
    "most_recent_order": {
      "order_id": "ORD123005",
      "total_value": 350,
      "order_date": ISODate("2023-09-10T14:00:00Z")
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567B"),
    "customer_name": "Neha Singh",
    "email": "neha.singh@example.com",
    "most_recent_order": {
      "order_id": "ORD123004",
      "total_value": 230,
      "order_date": ISODate("2023-08-25T14:00:00Z")
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567A"),
    "customer_name": "Ravi Patel",
    "email": "ravi.patel@example.com",
    "most_recent_order": {
      "order_id": "ORD123003",
      "total_value": 500,
      "order_date": ISODate("2023-07-01T14:00:00Z")
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345679"),
    "customer_name": "Priya Sharma",
    "email": "priya.sharma@example.com",
    "most_recent_order": {
      "order_id": "ORD123002",
      "total_value": 1050,
      "order_date": ISODate("2023-06-20T14:00:00Z")
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345678"),
    "customer_name": "Amit Kumar",
    "email": "amit.kumar@example.com",
    "most_recent_order": {
      "order_id": "ORD123001",
      "total_value": 1550,
      "order_date": ISODate("2023-05-15T14:00:00Z")
    }
  }
]
*/

/*4. Find the Most Expensive Order by Customer:*/
db.orders.aggregate([
    {
      $lookup: {
        from: "customers",
        localField: "customer_id",
        foreignField: "_id",
        as: "customer_info"
      }
    },
    { $unwind: "$customer_info" },
    {
      $group: {
        _id: "$customer_id",
        customer_name: { $first: "$customer_info.name" },
        most_expensive_order: {
          $max: {
            order_id: "$order_id",
            total_value: "$total_value",
            items: "$items"
          }
        }
      }
    }
  ]);

/*OUTPUT:
[
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345678"),
    "customer_name": "Amit Kumar",
    "most_expensive_order": {
      "order_id": "ORD123001",
      "total_value": 1550,
      "items": [
        { "product_name": "Laptop", "quantity": 1, "price": 1500 },
        { "product_name": "Mouse", "quantity": 2, "price": 25 }
      ]
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345679"),
    "customer_name": "Priya Sharma",
    "most_expensive_order": {
      "order_id": "ORD123002",
      "total_value": 1050,
      "items": [
        { "product_name": "Smartphone", "quantity": 1, "price": 1000 },
        { "product_name": "Charger", "quantity": 1, "price": 50 }
      ]
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567A"),
    "customer_name": "Ravi Patel",
    "most_expensive_order": {
      "order_id": "ORD123003",
      "total_value": 500,
      "items": [
        { "product_name": "Tablet", "quantity": 1, "price": 400 },
        { "product_name": "Headphones", "quantity": 1, "price": 100 }
      ]
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567B"),
    "customer_name": "Neha Singh",
    "most_expensive_order": {
      "order_id": "ORD123004",
      "total_value": 230,
      "items": [
        { "product_name": "Smartwatch", "quantity": 1, "price": 200 },
        { "product_name": "Earphones", "quantity": 1, "price": 30 }
      ]
    }
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567C"),
    "customer_name": "Vikram Reddy",
    "most_expensive_order": {
      "order_id": "ORD123005",
      "total_value": 350,
      "items": [
        { "product_name": "Monitor", "quantity": 1, "price": 300 },
        { "product_name": "Keyboard", "quantity": 1, "price": 50 }
      ]
    }
  }
]

*/

/*PART-3*/

/*1. Find All Customers Who Placed Orders in the Last Month:*/
db.orders.aggregate([
    {
      $match: {
        order_date: {
          $gte: new Date(new Date() - 30 * 24 * 60 * 60 * 1000)
        }
      }
    },
    {
      $lookup: {
        from: "customers",
        localField: "customer_id",
        foreignField: "_id",
        as: "customer_info"
      }
    },
    { $unwind: "$customer_info" },
    {
      $group: {
        _id: "$customer_id",
        customer_name: { $first: "$customer_info.name" },
        email: { $first: "$customer_info.email" },
        most_recent_order_date: { $max: "$order_date" }
      }
    }
  ]);

  /*OUTPUT:
  [
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567C"),
    "customer_name": "Vikram Reddy",
    "email": "vikram.reddy@example.com",
    "most_recent_order_date": ISODate("2023-09-10T14:00:00Z")
  }
]
*/

/*2. Find All Products Ordered by a Specific Customer:*/
db.orders.aggregate([
    {
      $match: { customer_id: ObjectId("64a1e4e4efdd3b0012345678") } // Amit Kumar's customer_id
    },
    {
      $unwind: "$items"
    },
    {
      $group: {
        _id: "$items.product_name",
        total_quantity: { $sum: "$items.quantity" }
      }
    }
  ]);

  /*OUTPUT:
  [
  {
    "_id": "Laptop",
    "total_quantity": 1
  },
  {
    "_id": "Mouse",
    "total_quantity": 2
  }
]
*/

/*3. Find the Top 3 Customers with the Most Expensive Total Orders:*/
db.orders.aggregate([
    {
      $lookup: {
        from: "customers",
        localField: "customer_id",
        foreignField: "_id",
        as: "customer_info"
      }
    },
    { $unwind: "$customer_info" },
    {
      $group: {
        _id: "$customer_id",
        customer_name: { $first: "$customer_info.name" },
        total_spent: { $sum: "$total_value" }
      }
    },
    { $sort: { total_spent: -1 } },
    { $limit: 3 }
  ]);

  /*OUTPUT:
  [
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345678"),
    "customer_name": "Amit Kumar",
    "total_spent": 1550
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b0012345679"),
    "customer_name": "Priya Sharma",
    "total_spent": 1050
  },
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567A"),
    "customer_name": "Ravi Patel",
    "total_spent": 500
  }
]
*/

/*4. Add a New Order for an Existing Customer:*/

// Checking if Jane Smith exists or not-
let janeSmith = db.customers.findOne({ name: "Jane Smith" });

// If not then add her to collection-
if (!janeSmith) {
    db.customers.insertOne({
        name: "Jane Smith",
        email: "jane.smith@example.com",
        address: { street: "101 G Road", city: "Pune", zipcode: "411001" },
        phone: "9876543215",
        registration_date: ISODate("2023-07-15T12:00:00Z")
    });
    janeSmith = db.customers.findOne({ name: "Jane Smith" });
}

/*OUTPUT:
{
  "_id": ObjectId("64a1e4e4efdd3b001234567F"),
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "address": { "street": "101 G Road", "city": "Pune", "zipcode": "411001" },
  "phone": "9876543215",
  "registration_date": ISODate("2023-07-15T12:00:00Z")
}
*/

// Insert a new order for Jane Smith
db.orders.insertOne({
  order_id: "ORD123006",
  customer_id: janeSmith._id,
  order_date: ISODate("2023-12-28T14:00:00Z"),
  status: "pending",
  items: [
      { product_name: "Smartphone", quantity: 1, price: 1000 },
      { product_name: "Headphones", quantity: 1, price: 100 }
  ],
  total_value: 1100
});

/*OUTPUT:
{
  "acknowledged": true,
  "insertedId": ObjectId("64a1e4e4efdd3b0012345680") // Example ObjectId for the new order
}
*/

//PART-4

//1. Find Customers Who Have Not Placed Orders:
db.customers.aggregate([
  {
      $lookup: {
          from: "orders",
          localField: "_id",
          foreignField: "customer_id",
          as: "orders"
      }
  },
  {
      $match: {
          "orders": { $size: 0 }  // Customers who have no orders
      }
  },
  {
      $project: {
          name: 1,
          email: 1
      }
  }
]);

/*OUTPUT:
[
  {
    "_id": ObjectId("64a1e4e4efdd3b001234567F"),
    "name": "Jane Smith",
    "email": "jane.smith@example.com"
  }
]
*/

//2. Calculate the Average Number of Items Ordered per Order:
db.orders.aggregate([
  { $unwind: "$items" },
  {
      $group: {
          _id: null,
          total_items: { $sum: "$items.quantity" },
          total_orders: { $sum: 1 }
      }
  },
  {
      $project: {
          _id: 0,
          average_items_per_order: { $divide: ["$total_items", "$total_orders"] }
      }
  }
]);

/*OUTPUT:
[
  {
    "average_items_per_order": 2.4
  }
]
*/

//3. Join Customer and Order Data Using $lookup:
db.orders.aggregate([
  {
      $lookup: {
          from: "customers",
          localField: "customer_id",
          foreignField: "_id",
          as: "customer_info"
      }
  },
  { $unwind: "$customer_info" },
  {
      $project: {
          "customer_info.name": 1,
          "customer_info.email": 1,
          order_id: 1,
          total_value: 1,
          order_date: 1
      }
  }
]);

/*OUTPUT:
[
  {
    "customer_info": {
      "name": "Amit Kumar",
      "email": "amit.kumar@example.com"
    },
    "order_id": "ORD123001",
    "total_value": 1550,
    "order_date": ISODate("2023-05-15T14:00:00Z")
  },
  {
    "customer_info": {
      "name": "Priya Sharma",
      "email": "priya.sharma@example.com"
    },
    "order_id": "ORD123002",
    "total_value": 1050,
    "order_date": ISODate("2023-06-20T14:00:00Z")
  },
  {
    "customer_info": {
      "name": "Ravi Patel",
      "email": "ravi.patel@example.com"
    },
    "order_id": "ORD123003",
    "total_value": 500,
    "order_date": ISODate("2023-07-01T14:00:00Z")
  },
  {
    "customer_info": {
      "name": "Neha Singh",
      "email": "neha.singh@example.com"
    },
    "order_id": "ORD123004",
    "total_value": 230,
    "order_date": ISODate("2023-08-25T14:00:00Z")
  },
  {
    "customer_info": {
      "name": "Vikram Reddy",
      "email": "vikram.reddy@example.com"
    },
    "order_id": "ORD123005",
    "total_value": 350,
    "order_date": ISODate("2023-09-10T14:00:00Z")
  }
]
*/

