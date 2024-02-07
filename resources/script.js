const input = document.querySelector(".input");
const send = document.querySelector(".send");
const update = document.querySelector(".update");
const list = document.querySelector(".list");

const newEP = "0.0.0.0:8080/new";
const statusEP = "0.0.0.0:8080/status";

send.addEventListener("click", async () => {
    const value = input.value;
    let resp = {
        ok: true,
        status: "404 Not Found",
        async text() {
            return 12;
        }
    };
    // let resp = await fetch(newEP, {
    //     method: "POST",
    //     body: value,
    //     headers: {
    //         "Content-Type": "text/plain"
    //     }
    // });
    await setTimeout(async () => {}, 2000);
    if (resp.ok) {
        //console.log("Waiting on response")
        let id = await resp.text();
        alert(`New expression added with id ${id}`);
        list.innerHTML += getItemHTML(id, value);
        input.value = "";
    } else {
        alert(`Expression could not be added. Status: ${resp.status}`);
    }
});

update.addEventListener("click", async () => {
    await updatelist(list)
})

//
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

async function updatelist(htmlList) {
    let resp = await fetch(statusEP, {
        method: "GET",
    })
    // json parsing here
    htmlList.innerHTML += getItemHTML())
}