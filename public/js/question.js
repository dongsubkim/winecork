function clickQ1() {
    console.log("Question 1 clicked!")
    const question = document.getElementById("q1-container")
    question.classList.toggle("active")
    let map = document.getElementById('map');
    map.classList.toggle("d-none")
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