### Start

0. Поставить Go если нет из официального источника [go.dev](https://go.dev/)

для линуксов:
```bash
# download the latest version
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz

# remove old version (if any)
sudo rm -rf /usr/local/go

# install the new version
sudo tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
```
для виндовсов - скачать файл установщика.

1. Добавить адреса для проверки в файл ./addresses/addressesSolana.txt . Один в строку.
2. Добавить API от [Alchemy](https://alchemy.com/?r=TA1MzI1MTg4NTc3N) в файл ./apis/alchemySolanaAPI.txt . Один в строку.
3. Компиляция и запуск
```bash
go clean
```
and
```bash
go build
```
запуск
```bash
./solist
```
для виндовс
```bash
./solist.exe
```
4. Результаты работы программы будут в ./balances/balancesSolana.txt
5. Благодарности:
Спасибо за бесплатное API: [https://www.alchemy.com/](https://alchemy.com/?r=TA1MzI1MTg4NTc3N)
6. Благодарю.



![Twitter Follow](https://img.shields.io/twitter/follow/huan_carlos?style=social)