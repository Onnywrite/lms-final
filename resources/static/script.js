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
    <button class="update-button text">Update</button>
</div>
`;
const expressionHtml = `
<div>
    <p class="no-content text">No content</p>
</div>
`;
let html = document.querySelector(".int-main");

let loadIndex = () => {
    html.innerHTML = indexHtml;
    openCalculatorPage();
};

let loadExpression = () => {
    html.innerHTML = expressionHtml;
    openExpressionsPage();
};


const calculatorPage = document.querySelector(".calculator-page");
const expressionsPage = document.querySelector(".expressions-page");

calculatorPage.addEventListener("click", () => loadIndex());
expressionsPage.addEventListener("click", () => loadExpression());

function openCalculatorPage() {
    const [expressionInput, additionTimeInput, subtractionTimeInput,
        multiplicationTimeInput, divisionTimeinpiut] = document.querySelectorAll(".expression-input");
    const send = document.querySelector(".send-button");
    const update = document.querySelector(".update-button");
    const list = document.querySelector(".list");

    const newEP = "http://localhost:8080/new";
    const statusEP = "http://localhost:8080/status";

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
            resp = JSON.parse(resp).id;
            alert(`New expression added with id ${resp}`);
            list.innerHTML += getItemHTML(id, expr);
            input.value = "";
        } else {
            alert(`Expression could not be added. Status: ${resp.status}`);
        }
    });

    update.addEventListener("click", async () => {
        let resp = await fetch(statusEP, {
            method: "GET",
        });
        // all temporary
        let obj = resp.json();
        htmlList.innerHTML += getItemHTML(obj.id, obj.expression);
    });

    // useless
    function getList(listOfElements) {
        listOfElements.forEach(item => {
            list.innerHTML += getItemHTML(item);
        })
    }

    const getItemHTML = (id, item) => `
    <li>
        ${id} - ${item}
    </li>`;
};

function openExpressionsPage() {

}
loadIndex();