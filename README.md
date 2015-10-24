## Demo: Extract, Transform, and Load (16m 55s)

Following along with this lesson on the Pluralsight course, I'm getting inconsistent results with the final code version.

Instead of the 4,000 odd records used in the online lesson, this demo only has three records.

I've followed the lession along and created three versions of the code Michael adds to the code.

First version has no channels, seconds has channels and the third uses shared channels.

I've checked the code several times but can't see any difference than how the tutorial video code explains.

The results are confusing. In the online video we only see the code run once, would it also have odd results if run multiple times?

**Step 1: No Channels (all good, three records processed)**

```
cd one
❯ go run main.go && cat dest.txt
13.690571ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            76502367              7       10.56       16.35          73.92         114.45
            56848544             18       18.51       21.68         333.18         390.24
            45687897              5       10.99       16.11          54.95          80.55
```

**Step 2: Add Channels (all good, three records processed)**

```
❯ cd ../two
❯ go run main.go && cat dest.txt
10.917755ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            76502367              7       10.56       16.35          73.92         114.45
            56848544             18       18.51       21.68         333.18         390.24
            45687897              5       10.99       16.11          54.95          80.55
```

**Step 3: Share channels across actors (sometimes three, sometimes two, sometimes one)**

```
❯ cd ../three
❯ go run main.go && cat dest.txt
4.809993ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55

❯ go run main.go && cat dest.txt
4.677514ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55
            56848544             18       18.51       21.68         333.18         390.24
            76502367              7       10.56       16.35          73.92         114.45
❯ go run main.go && cat dest.txt
4.733602ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55
            56848544             18       18.51       21.68         333.18         390.24
❯ go run main.go && cat dest.txt
4.75046ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            76502367              7       10.56       16.35          73.92         114.45
            45687897              5       10.99       16.11          54.95          80.55
❯ cd ../three
❯ go run main.go && cat dest.txt
4.809993ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55
❯ go run main.go && cat dest.txt
4.677514ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55
            56848544             18       18.51       21.68         333.18         390.24
            76502367              7       10.56       16.35          73.92         114.45
❯ go run main.go && cat dest.txt
4.733602ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            45687897              5       10.99       16.11          54.95          80.55
            56848544             18       18.51       21.68         333.18         390.24
❯ go run main.go && cat dest.txt
4.75046ms
         Part Number       Quantity   Unit Cost  Unit Price     Total Cost    Total Price
            76502367              7       10.56       16.35          73.92         114.45
            45687897              5       10.99       16.11          54.95          80.55
```
