# meecha_v5_back
meecha の バージョン5のリポジトリ

## 仕様技術
- タスク自動化: Taskfile
- コンテナ化: Docker
- 言語: Go

## クローンするとき
```
git clone https://github.com/meecha-org/meecha_v5_back --recursive 
```

## 環境構築
- 環境の構築
```
task setup
```

## 環境の削除
- 全部消す場合
```
docker compose down -v
```

- コンテナだけ消す場合
```
docker compose down
```

## コンテナのログの見方
```
docker compose logs -f app
```

## アクセスする方法
- フロントエンド
https://localhost:8633/ui/

- バックエンド
https://localhost:8633/app/

