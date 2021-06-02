function clickQ1() {
    console.log("Question 1 clicked!")
    const question = document.getElementById("q1-container");
    const pickUpLocation = document.querySelector("#pickUpLocation");
    const map = document.getElementById('map');
    console.log("marker clicked: ", markerClicked)
    if (markerClicked == true) {
        question.classList.remove("active")
    } else {
        question.classList.add("active")
        map.classList.remove("d-none")
    }
    markerClicked = false;
    // if (pickUpLocation.innerText.length > 0) {
    //     console.log("InnerText not empty")
    //     if (!question.classList.contains("active")) {
    //         question.classList.add("active")
    //     }
    // }
    // map.classList.toggle("d-none")
}

function clickQ2() {
    console.log("Question 2 clicked!")
    var question = document.querySelector(".question-2")
    var row = document.querySelector(".wine-selector-row")
    question.classList.toggle("active")
    row.classList.toggle("d-none")
}

function clickQ3() {
    console.log("Question 3 clicked!")
    let question = document.querySelector(".question-3")
    let row = document.querySelector(".price-selector")
    question.classList.toggle("active")
    row.classList.toggle("d-none")
}

function clickQ4() {
    console.log("Question 4 clicked!")
    let question = document.querySelector(".question-4")
    let row = document.querySelector(".food-matcher")
    question.classList.toggle("active")
    row.classList.toggle("d-none")
}

let wineType = document.getElementsByClassName("wine-selector");
Array.prototype.forEach.call(wineType, function (el) {
    el.addEventListener("click", wineTypeSelector)
})

function wineTypeSelector(e) {
    let wt = document.querySelector("#wineType");
    wt.innerText = this.innerText
    wt.classList.remove("d-none")
}

let priceSelector = document.getElementsByClassName("price-rect");

Array.prototype.forEach.call(priceSelector, function (el) {
    el.addEventListener("click", function () {
        let priceRange = document.querySelector("#priceRange");
        offPriceRange();
        this.classList.toggle("active");
        priceRange.innerText = this.innerText;
        priceRange.classList.remove("d-none");
    })
})

function offPriceRange() {
    let priceSelector = document.getElementsByClassName("price-rect");
    console.log(priceSelector);
    Array.prototype.forEach.call(priceSelector, function (el) {
        el.classList.remove("active")
    })
}

let foods = document.querySelectorAll(".food-match-food");
Array.prototype.forEach.call(foods, function (el) {
    el.addEventListener("click", function () {
        let foodMatched = document.querySelector("#foodMatch");
        foodMatched.classList.remove("d-none")
        foodMatched.innerText = this.innerText
        let btn = document.querySelector(".ok-btn")
        btn.classList.add("show")
    })
})

const storeSelector = document.querySelector(".store-selector");
const selectorBox = document.querySelector(".selector-box");
const selectOptions = document.querySelectorAll(".selector-row");
const selectBoxLine = document.querySelector(".selector-line");
const selectTitle = document.querySelector(".selector-row.title");
const selectedStore = document.querySelector(".selected-store");

storeSelector.addEventListener("click", function () {
    if (selectorBox.classList.contains("active")) {
        storeSelector.style = "z-index:1;"
        selectorBox.style = "z-index:1;"
        selectorBox.classList.toggle("active")
        selectOptions.forEach(element => {
            element.style = "display:none;"
        })
        selectBoxLine.style = "display:none;"
        selectedStore.style = "display:flex;"

    } else {
        storeSelector.style = "z-index:3;"
        selectorBox.style = "z-index:3;"
        selectorBox.classList.toggle("active")
        selectOptions.forEach(element => {
            element.style = "display:flex;"
        });
        selectBoxLine.style = "display:flex;"
        selectedStore.style = "display:none;"

    }
})


selectOptions.forEach(el => {
    el.addEventListener("click", function() {
        selectedStore.innerText = this.innerText
    })
})

const submitButton = document.querySelector(".ok-btn");
submitButton.addEventListener("click", function () {
    var url = "/";
    let store = document.querySelector("#pickUpLocation").innerText;
    let wine = document.querySelector("#wineType").innerText;
    let price = document.querySelector("#priceRange").innerText;
    let food = document.querySelector("#foodMatch").innerText;
    var params = `store=${store}&wine=${wine}&price=${price}&food=${food}`;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send(params);
})