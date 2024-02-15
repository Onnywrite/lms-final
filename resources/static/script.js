const newEP = "http://localhost:8081/new";
const statusEP = "http://localhost:8081/status";

const indexHtml = `
<input type="text" class="expression-input text" placeholder="Enter an expression">
<div class="expression-settings">
    <div>
        <h3 class="operations-tip text">Operations timings</h3>
    </div>
    <div class="operation-div">
        <p class="operation text">Addition</p>
        <input type="number" class="expression-input text" placeholder="Enter a number">
    </div>
    <div class="operation-div">
        <p class="operation text">Subtraction</p>
        <input type="number" class="expression-input text" placeholder="Enter a number">
    </div>
    <div class="operation-div">
        <p class="operation text">Multiplication</p>
        <input type="number" class="expression-input text" placeholder="Enter a number">
    </div>
    <div class="operation-div">
        <p class="operation text">Division</p>
        <input type="number" class="expression-input text" placeholder="Enter a number">
    </div>
</div>
<button class="send-button">Send</button>
<div>
    <ul class="list">

    </ul>
</div>
`;
const expressionsHtml = `
<div>
    <p>No content</p>
    <p>List of all expressions</p>
</div>
`;
const serversHtml = `
<div>
    <p>No content</p>
    <p>List of servers with their work status</p>
</div>
`;
const contactsHtml = `
<div>
    <p>No content</p>
    <p>contacts...</p>
</div>
`;
let html = document.querySelector(".int-main");

let loadIndex = () => {
    html.innerHTML = indexHtml;
    openCalculatorPage();
};

let loadExpressions = () => {
    html.innerHTML = expressionsHtml;
    openExpressionsPage();
};

let loadServers = () => {
    html.innerHTML = serversHtml;
    openServersPage();
};

let loadContacts = () => {
    html.innerHTML = contactsHtml;
    openContactsPage();
};

loadIndex();

const calculatorPage = document.querySelector(".calculator-page");
const expressionsPage = document.querySelector(".expressions-page");

calculatorPage.addEventListener("click", () => loadIndex());
expressionsPage.addEventListener("click", () => loadExpression());

function openCalculatorPage() {
    const [expressionInput, additionTimeInput, subtractionTimeInput,
        multiplicationTimeInput, divisionTimeinpiut] = document.querySelectorAll(".expression-input");
    const send = document.querySelector(".send-button");
    const list = document.querySelector(".list");
    
    let getItemHTML = (id, item) => `
    <li>
        ${id} - ${item}
    </li>`;

    send.addEventListener("click", async () => {
        const expr = {
            expression: expressionInput.value,
            addition_time: Number(additionTimeInput.value),
            subtraction_time: Number(subtractionTimeInput.value),
            multiplication_time: Number(multiplicationTimeInput.value),
            division_time: Number(divisionTimeinpiut.value),
        };
        let resp = await fetch(newEP, {
            method: "POST",
            body: JSON.stringify(expr),
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (resp.ok) {
            resp = await resp.text();
            alert(`New expression added with id ${resp}`);
            list.innerHTML += getItemHTML(0, expr);
        } else {
            alert(`Expression could not be added. Status: ${resp.status}`);
        }
    });

};

function openExpressionsPage() {

}

function openServersPage() {

}

function openContactsPage() {

}