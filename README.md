# feel-free-post-app

環境構築手順

1.リポジトリをクローン
```
git clone git@github.com:junmoro/feel-free-post-app.git
```
2.dockerイメージをビルド
```
cd 1でクローンしたディレクトリ
docker-compose build
```
3.docker起動
```
docker-compose up -d
```
4.react　→　http:localhost/
```
golang　→　http://localhost/api
phpMyAdmin　→　http://localhost:1234/
```
5.docker停止
```
docker-compose down
```
