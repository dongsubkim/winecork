const storeSelector = document.querySelector(".store-selector");
const selectorBox = document.querySelector(".selector-box");
const selectOptions = document.querySelectorAll(".selector-row");
const selectBoxLine = document.querySelector(".selector-line");
const selectTitle = document.querySelector(".selector-row.title");
const selectedStore = document.querySelector(".selected-store");

const submitForm = document.querySelector(".submit-form");
const submitButton = document.querySelector(".ok-btn");
const questionContainer = document.querySelector("#question-container");
const foodMatched = document.querySelector("#foodMatch");
const foodDetails = document.querySelectorAll(".food-match-detail");
const foodDetailCol = document.querySelectorAll(".food-match-detail-col");
let priceSelector = document.getElementsByClassName("price-rect");

let foodMatchSelected = false;
// questionContainer.addEventListener("click", allSelected)

function clickQ1() {
    let question = document.querySelector(".question-1")
    let row = document.querySelector(".food-matcher")
    if (!foodMatchSelected) {
        question.classList.add("active")
        row.classList.remove("d-none")
    } else {
        for (let row of foodDetails) {
            row.classList.remove("show");
        }
    }
    foodMatchSelected = false
}

function clickQ2() {
    if (answerFoodMatch.value.length == 0) {
        alert("음식을 먼저 골라주세요.")
        return
    }
    let question = document.querySelector(".question-2")
    let row = document.querySelector(".price-selector")
    question.classList.toggle("active")
    row.classList.toggle("d-none")
}

function clickQ3() {
    console.log("q3 clicked")
    if (answerPriceRange.value.length == 0) {
        alert("가격을 먼저 골라주세요.")
        return
    }
    const question = document.getElementById("q3-container");
    question.scrollIntoView({ behavior: "smooth", block: "end", inline: "nearest" })
    map.relayout();
    console.log("markerClicked:", markerClicked)
    if (markerClicked == true) {
        question.classList.remove("active")
        container.classList.add("d-none")
    } else {
        question.classList.add("active")
        container.classList.remove("d-none")
        storeSelector.classList.remove("d-none")
        show.classList.add("d-none")
        map.relayout();
    }
    markerClicked = false;
}

Array.prototype.forEach.call(priceSelector, function (el) {
    el.addEventListener("click", function () {
        let priceRange = document.querySelector("#priceRange");
        offPriceRange();
        this.classList.toggle("active");
        priceRange.innerText = this.innerText;
        answerPriceRange.value = this.id;
        priceRange.classList.remove("d-none");
    })
})

function offPriceRange() {
    let priceSelector = document.getElementsByClassName("price-rect");
    Array.prototype.forEach.call(priceSelector, function (el) {
        el.classList.remove("active")
    })
}

function foodDetailShow(foodType) {
    if (foodType == "salad" || foodType == "no-food") {
        let q = document.querySelector(".question-1");
        q.classList.remove("active");
        let row = document.querySelector(".food-matcher");
        row.classList.add("d-none")
        foodMatchSelected = true;
        return
    }
    for (let row of foodDetails) {
        let showRow = false;
        if (row.nodeName != "DIV") {
            continue
        }
        for (let col of row.childNodes) {
            if (col.nodeName != "DIV") {
                continue
            }
            if (col.classList.contains(foodType)) {
                showRow = true;
                col.classList.add("show");
            } else {
                col.classList.remove("show");
            }
        }
        if (showRow) {
            row.classList.add("show");
        } else {
            row.classList.remove("show");
        }
    }
}

Array.prototype.forEach.call(foodDetailCol, function (el) {
    el.addEventListener("click", function () {
        let i = foodMatched.innerText.indexOf(" | ");
        if (i > 0) {
            foodMatched.innerText = foodMatched.innerText.slice(0, i)
        }
        foodMatched.innerText = foodMatched.innerText.trim() + " | " + this.innerText
        answerFoodMatch.value = this.id
        let q = document.querySelector(".question-1");
        q.classList.remove("active");
        let row = document.querySelector(".food-matcher");
        row.classList.add("d-none")
        foodMatchSelected = true;
    })
})

let foods = document.querySelectorAll(".food-match-food");
Array.prototype.forEach.call(foods, function (el) {
    el.addEventListener("click", function () {
        foodMatched.classList.remove("d-none")
        foodMatched.innerText = this.innerText;
        answerFoodMatch.value = this.id;
        foodDetailShow(this.id)
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
        if (this.classList.contains("cand")) {
            for (let marker of markers) {
                storeType = marker.getTitle().split(" ")[0]
                if (storeType != this.innerText) {
                    marker.setVisible(false);
                } else {
                    marker.setVisible(true);
                }
            }
        }
    })
})

// function allSelected() {
//     if (answerStore && answerStore.value.length == 0) {
//         return false
//     }
//     if (answerPriceRange && answerPriceRange.value.length == 0) {
//         return false
//     }
//     if (answerFoodMatch && answerFoodMatch.value.length == 0) {
//         return false
//     }
//     submitForm.classList.remove("d-none")
//     submitButton.classList.add("show")
//     return true
// }
