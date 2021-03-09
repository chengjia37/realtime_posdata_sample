プログラム一覧
    gen_csv.go 
        CSVファイルを生成する。標準出力なのでファイルにリダイレクトして利用する
        shop_idにマイナスの値を指定するとレコードを出力する時間のsleep時間(ms)
        1shop=1プロセスの場合は実際の店舗でのCPU負荷やメモリを分析できる

    gen_csv.sh
        引数が多いgen_csv.goを起動するためのシェル
        windowsの場合、exeを実行


    read_csv.go
        gen_csvが出力したCSVファイルをポーリングしてKinesisに転送する
        指定が必要なのは対象ファイルのみ
        Kinesisの名前はposでハードコーディング


ファイル生成プログラムの実行
    $ ./gen_csv.sh
    $ tail -f data/before/20210216_SHOP-100_POS1.csv

CSVファイルの送信
    $ go run read_csv/read_csv.go data/before/20210216_SHOP-100_POS1.csv

Kinesisからの読み取り
    $ 