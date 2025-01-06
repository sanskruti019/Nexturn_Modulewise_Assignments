class ExpenseTracker {
    constructor() {
        this.expenses = JSON.parse(localStorage.getItem('expenses')) || [];
        this.form = document.getElementById('expenseForm');
        this.expensesList = document.getElementById('expensesList');
        this.chart = null;

        this.categoryColors = {
            Food: '#3b82f6',
            Travel: '#10b981',
            Shopping: '#f59e0b',
            Entertainment: '#8b5cf6',
            Bills: '#ef4444'
        };

        this.initializeEventListeners();
        this.renderExpenses();
        this.updateChart();
    }

    initializeEventListeners() {
        this.form.addEventListener('submit', (e) => {
            e.preventDefault();
            this.addExpense();
        });
    }

    addExpense() {
        const expense = {
            id: Date.now(),
            amount: parseFloat(document.getElementById('amount').value),
            description: document.getElementById('description').value,
            category: document.getElementById('category').value,
            date: new Date().toLocaleDateString()
        };

        this.expenses.push(expense);
        this.saveExpenses();
        this.renderExpenses();
        this.updateChart();
        this.form.reset();
    }

    deleteExpense(id) {
        this.expenses = this.expenses.filter(expense => expense.id !== id);
        this.saveExpenses();
        this.renderExpenses();
        this.updateChart();
    }

    saveExpenses() {
        localStorage.setItem('expenses', JSON.stringify(this.expenses));
    }

    renderExpenses() {
        this.expensesList.innerHTML = '';
        this.expenses.forEach(expense => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${expense.date}</td>
                <td>${expense.description}</td>
                <td>
                    <span class="category-badge" style="background-color: ${this.categoryColors[expense.category]}20; color: ${this.categoryColors[expense.category]}">
                        ${expense.category}
                    </span>
                </td>
                <td>$${expense.amount.toFixed(2)}</td>
                <td>
                    <button class="delete-btn" onclick="expenseTracker.deleteExpense(${expense.id})">Delete</button>
                </td>
            `;
            this.expensesList.appendChild(row);
        });

        this.renderCategoryTotals();
    }

    renderCategoryTotals() {
        const totals = this.getCategoryTotals();
        const categoryTotalsDiv = document.getElementById('categoryTotals');
        categoryTotalsDiv.innerHTML = Object.entries(totals)
            .map(([category, total]) => `
                <div style="display: flex; justify-content: space-between; margin: 0.5rem 0;">
                    <span style="color: ${this.categoryColors[category]}">${category}</span>
                    <strong>$${total.toFixed(2)}</strong>
                </div>
            `)
            .join('');
    }

    getCategoryTotals() {
        return this.expenses.reduce((totals, expense) => {
            totals[expense.category] = (totals[expense.category] || 0) + expense.amount;
            return totals;
        }, {});
    }

    updateChart() {
        const totals = this.getCategoryTotals();
        const ctx = document.getElementById('expenseChart').getContext('2d');

        if (this.chart) {
            this.chart.destroy();
        }

        this.chart = new Chart(ctx, {
            type: 'doughnut',
            data: {
                labels: Object.keys(totals),
                datasets: [{
                    data: Object.values(totals),
                    backgroundColor: Object.keys(totals).map(category => this.categoryColors[category]),
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: {
                        position: 'bottom'
                    },
                    title: {
                        display: true,
                        text: 'Expenses by Category'
                    }
                }
            }
        });
    }
}

const expenseTracker = new ExpenseTracker();