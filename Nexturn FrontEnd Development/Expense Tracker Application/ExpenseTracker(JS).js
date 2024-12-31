let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

function updateCategorySummary() {
    let foodTotal = 0, travelTotal = 0, shoppingTotal = 0;

    for (let expense of expenses) {
        if (expense.category === "Food") foodTotal += expense.amount;
        if (expense.category === "Travel") travelTotal += expense.amount;
        if (expense.category === "Shopping") shoppingTotal += expense.amount;
    }

    document.querySelector("#categorySummary").innerHTML = `
        <li>Food: ₹${foodTotal}</li>
        <li>Travel: ₹${travelTotal}</li>
        <li>Shopping: ₹${shoppingTotal}</li>
    `;
}

function addExpense() {
    const amount = parseFloat(document.getElementById('amount').value);
    const description = document.getElementById('description').value;
    const category = document.getElementById('category').value;

    if (isNaN(amount) || description === "") {
        alert("Please enter valid values!");
        return;
    }

    const expense = { amount, description, category };

    expenses.push(expense);
    localStorage.setItem('expenses', JSON.stringify(expenses));

    updateCategorySummary();
    renderExpenses();
    resetForm();
}

function renderExpenses() {
    const expenseList = document.getElementById('expenseList');
    expenseList.innerHTML = "";

    expenses.forEach((expense, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>₹${expense.amount}</td> 
            <td>${expense.description}</td> 
            <td>${expense.category}</td>
            <td><button onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseList.appendChild(row);
    });
}

function deleteExpense(index) {
    expenses.splice(index, 1);
    localStorage.setItem('expenses', JSON.stringify(expenses));
    updateCategorySummary();
    renderExpenses();
}

function resetForm() {
    document.getElementById('amount').value = "";
    document.getElementById('description').value = "";
    document.getElementById('category').value = "Food";
}

function init() {
    renderExpenses();
    updateCategorySummary();
}

init();
