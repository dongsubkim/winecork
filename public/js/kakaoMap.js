const show = document.querySelector("#pickUpLocation");
const question = document.querySelector(".question-1");
const store = document.querySelector(".store-selector");

const answerStore = document.querySelector("#answerStore");
const answerWineType = document.querySelector("#answerWineType");
const answerPriceRange = document.querySelector("#answerPriceRange");
const answerFoodMatch = document.querySelector("#answerFoodMatch");

var container = document.getElementById('map'); //지도를 담을 영역의 DOM 레퍼런스

var options = { //지도를 생성할 때 필요한 기본 옵션
    center: new kakao.maps.LatLng(37.28734256346641, 127.0596781925285), //지도의 중심좌표.
    level: 5 //지도의 레벨(확대, 축소 정도)
};
var map = new kakao.maps.Map(container, options); //지도 생성 및 객체 리턴

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition((pos) => {
            map.setCenter(new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude))
        });
    } else {
        console.log("Geolocation is not supported by this browser.")
    }
}
getLocation()

const lotteMart = [
    {
        store: "롯데마트",
        lat: 37.28734256346641,
        lng: 127.0596781925285
    }
]

const lotteDep = [
    {
        store: "롯데백화점",
        location: "본점",
        lat: 37.28834256346641,
        lng: 127.0596781925285
    }
]

const emart = [
    {
        store: "이마트",
        location: "가양점",
        lat: 37.28934256346641,
        lng: 127.0596781925285
    }
]

const ssg = [
    {
        store: "신세계백화점",
        location: "본점",
        lat: 37.28634256346641,
        lng: 127.0596781925285
    }
]

function createMarker(info) {
    let markerPos = new kakao.maps.LatLng(info.lat, info.lng);
    let title = info.store
    if (info.store != "롯데마트") {
        title += " " + info.location
    }
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

for (let store of lotteMart) {
    createMarker(store)
}

for (let store of lotteDep) {
    createMarker(store)
}

for (let store of emart) {
    createMarker(store)
}

for (let store of ssg) {
    createMarker(store)
}



// // 마커가 표시될 위치입니다 
// var markerPosition1 = new kakao.maps.LatLng(33.450701, 126.570667);
// var markerPosition2 = new kakao.maps.LatLng(33.451701, 126.570667);
// var markerPosition3 = new kakao.maps.LatLng(33.452701, 126.570667);
// var markerPosition4 = new kakao.maps.LatLng(33.453701, 126.570667);

// // 마커를 생성합니다
// var marker1 = new kakao.maps.Marker({
//     position: markerPosition1,
//     title: "롯데마트",
//     clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
// });
// // 마커를 생성합니다
// var marker2 = new kakao.maps.Marker({
//     position: markerPosition2,
//     title: "롯데백화점 본점",
//     clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
// });

// // 마커를 생성합니다
// var marker3 = new kakao.maps.Marker({
//     position: markerPosition3,
//     title: "이마트몰 가양점",
//     clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
// });

// // 마커를 생성합니다
// var marker4 = new kakao.maps.Marker({
//     position: markerPosition4,
//     title: "신세계백화점 강남점",
//     clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
// });


// // 마커가 지도 위에 표시되도록 설정합니다
// marker1.setMap(map);
// marker2.setMap(map);
// marker3.setMap(map);
// marker4.setMap(map);

// // 마커에 클릭이벤트를 등록합니다
// kakao.maps.event.addListener(marker1, 'click', clickMarker);
// kakao.maps.event.addListener(marker2, 'click', clickMarker);
// kakao.maps.event.addListener(marker3, 'click', clickMarker);
// kakao.maps.event.addListener(marker4, 'click', clickMarker);
