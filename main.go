package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Sunucu durumunu simüle eden rastgele veriler üreten fonksiyon
func rastgeleMetrikUret() (int, int) {
	// Sayfa her yenilendiğinde %40 ile %95 arasında rastgele CPU ve RAM kullanımı uydurur
	cpu := rand.Intn(55) + 40
	ram := rand.Intn(45) + 50
	return cpu, ram
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	cpu, ram := rastgeleMetrikUret()
	simdi := time.Now().Format("15:04:05 - 02/01/2006")

	// Tarayıcıya HTML (görsel içerik) göndereceğimizi belirtiyoruz
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Şık bir HTML ve CSS arayüzü basıyoruz
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Biner Cloud Dashboard</title>
		<meta http-equiv="refresh" content="1"> <style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #1a1a1a; color: #ffffff; text-align: center; padding-top: 50px; }
			.container { background-color: #2d2d2d; display: inline-block; padding: 30px; border-radius: 15px; box-shadow: 0px 4px 15px rgba(0,0,0,0.5); }
			h1 { color: #00ffd2; margin-bottom: 5px; }
			.subtitle { color: #aaa; font-size: 14px; margin-bottom: 25px; }
			.metric-box { display: inline-block; width: 150px; background: #222; padding: 15px; margin: 10px; border-radius: 8px; border: 1px solid #444; }
			.metric-value { font-size: 24px; font-weight: bold; color: #ff007f; }
			.live-indicator { display: inline-block; width: 10px; height: 10px; background-color: #00ff00; border-radius: 50px; margin-right: 5px; animation: blink 1s infinite; }
			@keyframes blink { 0%% { opacity: 0.2; } 50%% { opacity: 1; } 100%% { opacity: 0.2; } }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Biner Cloud Dashboard v3</h1>
			<div class="subtitle">Sistem Altyapısı Tamamen Go ile Yönetiliyor</div>
			
			<div class="metric-box">
				<div>CPU Kullanımı</div>
				<div class="metric-value">%%%d</div>
			</div>
			
			<div class="metric-box">
				<div>RAM Kullanımı</div>
				<div class="metric-value">%%%d</div>
			</div>
			
			<div style="margin-top: 20px; font-size: 14px; color: #00ffd2;">
				<span class="live-indicator"></span> Sunucu Zamanı: %s
			</div>
			<p style="font-size: 12px; color: #666;">F5 yapmana gerek yok, sayfa 3 saniyede bir canlı güncellenir abi.</p>
		</div>
	</body>
	</html>
	`, cpu, ram, simdi)
}

func main() {
	// Go 1.22+ ve 1.25 sürümlerinde rand.Seed kullanmaya gerek yoktur, otomatik tetiklenir.
	http.HandleFunc("/", dashboardHandler)
	fmt.Println("Aksiyonlu Dashboard 8080 portunda baslatıldı...")
	http.ListenAndServe(":8080", nil)
}
