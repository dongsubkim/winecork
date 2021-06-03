const storeSelector = document.querySelector(".store-selector");
const selectorBox = document.querySelector(".selector-box");
const selectOptions = document.querySelectorAll(".selector-row");
const selectBoxLine = document.querySelector(".selector-line");
const selectTitle = document.querySelector(".selector-row.title");
const selectedStore = document.querySelector(".selected-store");

const submitForm = document.querySelector(".submit-form");
const submitButton = document.querySelector(".ok-btn");
const questionContainer = document.querySelector("#question-container");
questionContainer.addEventListener("click", allSelected)

function clickQ1() {
    console.log("Question 1 clicked!")
    const question = document.getElementById("q1-container");
    const map = document.getElementById('map');
    console.log("marker clicked: ", markerClicked)
    if (markerClicked == true) {
        question.classList.remove("active")
    } else {
        question.classList.add("active")
        map.classList.remove("d-none")
    }
    markerClicked = false;
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
    wt.innerText = this.innerText;
    wt.classList.remove("d-none");
    answerWineType.value = this.innerText;
}

let priceSelector = document.getElementsByClassName("price-rect");

Array.prototype.forEach.call(priceSelector, function (el) {
    el.addEventListener("click", function () {
        let priceRange = document.querySelector("#priceRange");
        offPriceRange();
        this.classList.toggle("active");
        priceRange.innerText = this.innerText;
        answerPriceRange.value = this.innerText;
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
        foodMatched.innerText = this.innerText;
        answerFoodMatch.value = this.innerText;
    })
})


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
    el.addEventListener("click", function () {
        selectedStore.innerText = this.innerText
    })
})

function allSelected() {
    if (answerStore && answerStore.value.length == 0) {
        return false
    }
    if (answerPriceRange && answerPriceRange.value.length == 0) {
        return false
    }
    if (answerWineType && answerWineType.value.length == 0) {
        return false
    }
    if (answerFoodMatch && answerFoodMatch.value.length == 0) {
        return false
    }
    submitForm.classList.remove("d-none")
    submitButton.classList.add("show")
    return true
}

