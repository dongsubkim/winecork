const show = document.querySelector("#pickUpLocation");
const question = document.querySelector(".question-1");
const store = document.querySelector(".store-selector");

const answerStore = document.querySelector("#answerStore");
const answerWineType = document.querySelector("#answerWineType");
const answerPriceRange = document.querySelector("#answerPriceRange");
const answerFoodMatch = document.querySelector("#answerFoodMatch");

var container = document.getElementById('map'); //지도를 담을 영역의 DOM 레퍼런스
var options = { //지도를 생성할 때 필요한 기본 옵션
    center: new kakao.maps.LatLng(33.450701, 126.570667), //지도의 중심좌표.
    level: 5 //지도의 레벨(확대, 축소 정도)
};

var map = new kakao.maps.Map(container, options); //지도 생성 및 객체 리턴

// 마커가 표시될 위치입니다 
var markerPosition1 = new kakao.maps.LatLng(33.450701, 126.570667);
var markerPosition2 = new kakao.maps.LatLng(33.451701, 126.570667);
var markerPosition3 = new kakao.maps.LatLng(33.452701, 126.570667);
var markerPosition4 = new kakao.maps.LatLng(33.453701, 126.570667);

// 마커를 생성합니다
var marker1 = new kakao.maps.Marker({
    position: markerPosition1,
    title: "롯데마트 1",
    clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
});
// 마커를 생성합니다
var marker2 = new kakao.maps.Marker({
    position: markerPosition2,
    title: "롯데백화점 1",
    clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
});

// 마커를 생성합니다
var marker3 = new kakao.maps.Marker({
    position: markerPosition3,
    title: "이마트 1",
    clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
});

// 마커를 생성합니다
var marker4 = new kakao.maps.Marker({
    position: markerPosition4,
    title: "신세계백화점 1",
    clickable: true // 마커를 클릭했을 때 지도의 클릭 이벤트가 발생하지 않도록 설정합니다
});


// 마커가 지도 위에 표시되도록 설정합니다
marker1.setMap(map);
marker2.setMap(map);
marker3.setMap(map);
marker4.setMap(map);

function clickMarker() {
    show.innerText = this.getTitle();
    answerStore.value = this.getTitle();
    show.classList.remove("d-none")
    container.classList.add("d-none")
    store.classList.add("d-none")
    question.classList.remove("active")
    markerClicked = true
}

// 마커에 클릭이벤트를 등록합니다
kakao.maps.event.addListener(marker1, 'click', clickMarker);
kakao.maps.event.addListener(marker2, 'click', clickMarker);
kakao.maps.event.addListener(marker3, 'click', clickMarker);
kakao.maps.event.addListener(marker4, 'click', clickMarker);
