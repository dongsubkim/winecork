const show = document.querySelector("#pickUpLocation");
const question = document.querySelector(".question-3");
const store = document.querySelector(".store-selector");

const answerStore = document.querySelector("#answerStore");
const answerPriceRange = document.querySelector("#answerPriceRange");
const answerFoodMatch = document.querySelector("#answerFoodMatch");

const stores = ["롯데마트", "롯데백화점", "신세계백화점", "이마트"]

var container = document.getElementById('map'); //지도를 담을 영역의 DOM 레퍼런스

var markers = [];

var map = new kakao.maps.Map(container, { //지도를 생성할 때 필요한 기본 옵션
    center: new kakao.maps.LatLng(37.28734256346641, 127.0596781925285), //지도의 중심좌표.
    level: 6 //지도의 레벨(확대, 축소 정도)
}); //지도 생성 및 객체 리턴

getLocation()

for (let store of storeLocations) {
    createMarker(store)
}

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition((pos) => {
            map.setCenter(new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude));
            map.relayout();
            map.setLevel(8);
            map.relayout();
            // for (let store of stores) {
            //     places.keywordSearch(store, searchCallback, {
            //         useMapCenter: true,
            //         radius: 10000
            //     });
            // }
        });
    }
}

function createMarker(info) {
    if (typeof info.Latitude == 'number') {
        var markerPos = new kakao.maps.LatLng(info.Latitude, info.Longitude);
    } else {
        var markerPos = new kakao.maps.LatLng(parseFloat(info.Latitude), parseFloat(info.Longitude));
    }

    let title = info.StoreType + " " + info.Location
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
    submitForm.classList.remove("d-none")
    submitButton.classList.add("show")
}

// var places = new kakao.maps.services.Places(map);
// var searchCallback = function (result, status, pagination) {
//     if (status === kakao.maps.services.Status.OK) {
//         for (let store of result) {
//             names = store.place_name.split(" ");
//             if (!storeLocations.includes(store.place_name) && names[0] != "롯데마트") {
//                 continue
//             }
//             if (stores.includes(names[0]) && names.length == 2) {
//                 info = {
//                     StoreType: names[0],
//                     Location: names[1],
//                     Latitude: parseFloat(store.y),
//                     Longitude: parseFloat(store.x)
//                 }
//                 createMarker(info)
//             }
//         }
//         if (pagination.hasNextPage) {
//             pagination.nextPage()
//         }
//     }
// };
