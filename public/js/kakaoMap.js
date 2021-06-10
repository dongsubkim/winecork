const show = document.querySelector("#pickUpLocation");
const question = document.querySelector(".question-1");
const store = document.querySelector(".store-selector");

const answerStore = document.querySelector("#answerStore");
const answerWineType = document.querySelector("#answerWineType");
const answerPriceRange = document.querySelector("#answerPriceRange");
const answerFoodMatch = document.querySelector("#answerFoodMatch");

const stores = ["롯데마트", "롯데백화점", "신세계백화점", "이마트"]

var container = document.getElementById('map'); //지도를 담을 영역의 DOM 레퍼런스

var options = { //지도를 생성할 때 필요한 기본 옵션
    center: new kakao.maps.LatLng(37.28734256346641, 127.0596781925285), //지도의 중심좌표.
    level: 5 //지도의 레벨(확대, 축소 정도)
};
var map = new kakao.maps.Map(container, options); //지도 생성 및 객체 리턴

var markers = [];

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition((pos) => {
            map.setCenter(new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude))
            for (let store of stores) {
                places.keywordSearch(store, searchCallback, {
                    useMapCenter: true,
                    radius: 10000
                });
            }
        });
    } else {
        console.log("Geolocation is not supported by this browser.")
    }
}
getLocation()

// const lotteMart = [
//     {
//         store: "롯데마트",
//         lat: 37.28734256346641,
//         lng: 127.0596781925285
//     }
// ]

// const lotteDep = [
//     {
//         store: "롯데백화점",
//         location: "본점",
//         lat: 37.28834256346641,
//         lng: 127.0596781925285
//     }
// ]

// const emart = [
//     {
//         store: "이마트",
//         location: "가양점",
//         lat: 37.28934256346641,
//         lng: 127.0596781925285
//     }
// ]

// const ssg = [
//     {
//         store: "신세계백화점",
//         location: "본점",
//         lat: 37.28634256346641,
//         lng: 127.0596781925285
//     }
// ]

function createMarker(info) {
    let markerPos = new kakao.maps.LatLng(info.lat, info.lng);
    let title = info.store + " " + info.location
    let marker = new kakao.maps.Marker({
        position: markerPos,
        title: title,
        clickable: true,
        map: map,
    })
    kakao.maps.event.addListener(marker, 'click', clickMarker);
    // 마커에 커서가 오버됐을 때 마커 위에 표시할 인포윈도우를 생성합니다
    var iwContent = `<div class="info-window" style="border-radius:9px;">${title}</div>`; // 인포윈도우에 표출될 내용으로 HTML 문자열이나 document element가 가능합니다

    // 인포윈도우를 생성합니다
    var infowindow = new kakao.maps.InfoWindow({
        content: iwContent
    });

    // 마커에 마우스오버 이벤트를 등록합니다
    kakao.maps.event.addListener(marker, 'mouseover', function () {
        // 마커에 마우스오버 이벤트가 발생하면 인포윈도우를 마커위에 표시합니다
        infowindow.open(map, marker);
    });

    // 마커에 마우스아웃 이벤트를 등록합니다
    kakao.maps.event.addListener(marker, 'mouseout', function () {
        // 마커에 마우스아웃 이벤트가 발생하면 인포윈도우를 제거합니다
        infowindow.close();
    });
    markers.push(marker);
}


function clickMarker() {
    show.innerText = this.getTitle();
    answerStore.value = this.getTitle();
    show.classList.remove("d-none")
    container.classList.add("d-none")
    store.classList.add("d-none")
    question.classList.remove("active")
    markerClicked = true
}

// for (let store of lotteMart) {
//     createMarker(store)
// }

// for (let store of lotteDep) {
//     createMarker(store)
// }

// for (let store of emart) {
//     createMarker(store)
// }

// for (let store of ssg) {
//     createMarker(store)
// }

var places = new kakao.maps.services.Places(map);
var searchCallback = function (result, status, pagination) {
    if (status === kakao.maps.services.Status.OK) {
        for (let store of result) {
            names = store.place_name.split(" ");
            if (!storeLocations.includes(store.place_name) && names[0] != "롯데마트") {
                continue
            }
            if (stores.includes(names[0]) && names.length == 2) {
                info = {
                    store: names[0],
                    location: names[1],
                    lat: parseFloat(store.y),
                    lng: parseFloat(store.x)
                }
                createMarker(info)
            }
        }
        if (pagination.hasNextPage) {
            pagination.nextPage()
        }
    }
};

