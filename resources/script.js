const input = document.querySelector(".input");
const send = document.querySelector(".send");
const update = document.querySelector(".update");
const list = document.querySelector(".list");

const indexEP = "0.0.0.0:8080/";
const newEP = "0.0.0.0:8080/new";
const statusEP = "0.0.0.0:8080/status";

send.addEventListener("click", async () => {
    const expr = {expression: input.value };
    let resp = await fetch(newEP, {
        method: "POST",
        body: JSON.stringify(expr),
        headers: {
            "Content-Type": "application/json"
        }
    });
    if (resp.ok) {
        let id = await resp.json().id;
        alert(`New expression added with id ${id}`);
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
    let obj = resp.json()
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
    </li>

`