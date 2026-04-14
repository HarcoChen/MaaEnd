# AutoStockpile → BuyGoods / BuyGoodsElastic 完整更名计划

## 规则总览

- **Go-service**：统一更名为 `buygoodselastic`，因为当前 Go 逻辑仅服务于 Elastic 流程。
- **Pipeline**：严格按照所在文件夹确定前缀。
    - `Main.json`：使用总前缀 `BuyGoods`
    - `Elastic/*`：使用前缀 `BuyGoodsElastic`
    - `Staple/*`：使用前缀 `BuyGoodsStaple`
- **节点名中若含有 `Staple` / `Elastic`**：将语义并入新前缀，避免重复词。
    - 例如：
        - `AutoStockpileInStapleItem` → `BuyGoodsStapleInItem`
        - `AutoStockpileElasticGoodsButton` → `BuyGoodsElasticGoodsButton`

---

## 一、Go-service 关键更名计划

### 1.1 包与目录

| 文件/范围                             | 类型        | 当前名称                   | 目标名称                     | 说明              |
| ------------------------------------- | ----------- | -------------------------- | ---------------------------- | ----------------- |
| `agent/go-service/autostockpile/`     | directory   | `autostockpile`            | `buygoodselastic`            | Go 包目录统一更名 |
| `agent/go-service/autostockpile/*.go` | package     | `package autostockpile`    | `package buygoodselastic`    | 全包统一          |
| `agent/go-service/register.go`        | import/call | `autostockpile.Register()` | `buygoodselastic.Register()` | 注册入口同步更新  |

### 1.2 Go 与 Pipeline 的契约名

| 文件                                      | 类型               | 当前名称                          | 目标名称                            | 说明                                       |
| ----------------------------------------- | ------------------ | --------------------------------- | ----------------------------------- | ------------------------------------------ |
| `agent/go-service/autostockpile/nodes.go` | custom action      | `AutoStockpile.SelectItem`        | `BuyGoodsElastic.SelectItem`        | `custom_action` 契约，必须与 pipeline 同步 |
| `agent/go-service/autostockpile/nodes.go` | custom action      | `AutoStockpile.ReconcileDecision` | `BuyGoodsElastic.ReconcileDecision` | 同上                                       |
| `agent/go-service/autostockpile/nodes.go` | custom recognition | `AutoStockpile.Recognition`       | `BuyGoodsElastic.Recognition`       | `custom_recognition` 契约                  |

### 1.3 Go 中直接引用的 Pipeline 节点常量

| 文件                                      | 类型        | 当前名称                              | 目标名称                                | 说明                      |
| ----------------------------------------- | ----------- | ------------------------------------- | --------------------------------------- | ------------------------- |
| `agent/go-service/autostockpile/nodes.go` | const       | `autoStockpileComponent`              | `buyGoodsElasticComponent`              | 日志组件名建议同步        |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileRelayNodeDecisionReady` | `BuyGoodsElasticRelayNodeDecisionReady` | Go 直接引用 pipeline 节点 |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileSelectedGoodsClick`     | `BuyGoodsElasticSelectedGoodsClick`     | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileSwipeMax`               | `BuyGoodsElasticSwipeMax`               | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileSwipeSpecificQuantity`  | `BuyGoodsElasticSwipeSpecificQuantity`  | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileSkip`                   | `BuyGoodsElasticSkip`                   | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileAttach`                 | `BuyGoodsElasticAttach`                 | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileFindMarketMark`         | `BuyGoodsElasticFindMarketMark`         | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileGetQuota`               | `BuyGoodsElasticGetQuota`               | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileGetQuotaAddition`       | `BuyGoodsElasticGetQuotaAddition`       | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileLocateGoods`            | `BuyGoodsElasticLocateGoods`            | 同上                      |
| `agent/go-service/autostockpile/nodes.go` | const value | `AutoStockpileGetGoods`               | `BuyGoodsElasticGetGoods`               | 同上                      |

---

## 二、Pipeline 全量节点更名计划

### 2.1 `assets/resource/pipeline/AutoStockpile/Main.json`

| 当前节点名                     | 目标节点名                |
| ------------------------------ | ------------------------- |
| `AutoStockpileMain`            | `BuyGoodsMain`            |
| `AutoStockpileElasticValleyIV` | `BuyGoodsElasticValleyIV` |
| `AutoStockpileStapleValleyIV`  | `BuyGoodsStapleValleyIV`  |
| `AutoStockpileElasticWuling`   | `BuyGoodsElasticWuling`   |
| `AutoStockpileStapleWuling`    | `BuyGoodsStapleWuling`    |
| `AutoStockpileDone`            | `BuyGoodsDone`            |

### 2.2 `assets/resource/pipeline/AutoStockpile/Elastic/Entry.json`

| 当前节点名                          | 目标节点名                         |
| ----------------------------------- | ---------------------------------- |
| `AutoStockpileElasticGoodsButton`   | `BuyGoodsElasticGoodsButton`       |
| `AutoStockpileSelectMarket`         | `BuyGoodsElasticSelectMarket`      |
| `AutoStockpileTask`                 | `BuyGoodsElasticTask`              |
| `AutoStockpileCheckInStore`         | `BuyGoodsElasticCheckInStore`      |
| `AutoStockpileGotoElasticGoods`     | `BuyGoodsElasticGotoGoods`         |
| `AutoStockpileFindMarketSidebar`    | `BuyGoodsElasticFindMarketSidebar` |
| `AutoStockpileElasticGoods`         | `BuyGoodsElasticGoods`             |
| `AutoStockpileElasticGoodsSelected` | `BuyGoodsElasticGoodsSelected`     |
| `AutoStockpileEnaureElasticClicked` | `BuyGoodsElasticEnaureClicked`     |

### 2.3 `assets/resource/pipeline/AutoStockpile/Elastic/DecisionLoop.json`

| 当前节点名                             | 目标节点名                               |
| -------------------------------------- | ---------------------------------------- |
| `AutoStockpileDecisionWuling`          | `BuyGoodsElasticDecisionWuling`          |
| `AutoStockpileDecisionValleyIV`        | `BuyGoodsElasticDecisionValleyIV`        |
| `AutoStockpileClearDecisionHitCount`   | `BuyGoodsElasticClearDecisionHitCount`   |
| `AutoStockpileRelayNodeDecision`       | `BuyGoodsElasticRelayNodeDecision`       |
| `AutoStockpileSelectedGoodsClick`      | `BuyGoodsElasticSelectedGoodsClick`      |
| `AutoStockpileEnsureDetailFrameStable` | `BuyGoodsElasticEnsureDetailFrameStable` |
| `AutoStockpileEnsureCashAnchor`        | `BuyGoodsElasticEnsureCashAnchor`        |
| `AutoStockpileBackToSelection`         | `BuyGoodsElasticBackToSelection`         |
| `AutoStockpileSkip`                    | `BuyGoodsElasticSkip`                    |
| `AutoStockpileDecisionFailed`          | `BuyGoodsElasticDecisionFailed`          |
| `AutoStockpileReconcileDecision`       | `BuyGoodsElasticReconcileDecision`       |
| `AutoStockpileAttach`                  | `BuyGoodsElasticAttach`                  |

### 2.4 `assets/resource/pipeline/AutoStockpile/Elastic/Helper.json`

| 当前节点名                      | 目标节点名                        |
| ------------------------------- | --------------------------------- |
| `AutoStockpileGetGoods`         | `BuyGoodsElasticGetGoods`         |
| `AutoStockpileLocateGoods`      | `BuyGoodsElasticLocateGoods`      |
| `AutoStockpileGoodsFilter`      | `BuyGoodsElasticGoodsFilter`      |
| `AutoStockpileGetQuotaAddition` | `BuyGoodsElasticGetQuotaAddition` |
| `AutoStockpileGetQuota`         | `BuyGoodsElasticGetQuota`         |
| `AutoStockpileFindMarketMark`   | `BuyGoodsElasticFindMarketMark`   |
| `AutoStockpileGetStockBill`     | `BuyGoodsElasticGetStockBill`     |

### 2.5 `assets/resource/pipeline/AutoStockpile/Elastic/Purchase.json`

| 当前节点名                            | 目标节点名                              |
| ------------------------------------- | --------------------------------------- |
| `AutoStockpileRelayNodeDecisionReady` | `BuyGoodsElasticRelayNodeDecisionReady` |
| `AutoStockpileCheckBuy`               | `BuyGoodsElasticCheckBuy`               |
| `AutoStockpileCheckQuota`             | `BuyGoodsElasticCheckQuota`             |
| `AutoStockpileSwipeSpecificQuantity`  | `BuyGoodsElasticSwipeSpecificQuantity`  |
| `AutoStockpileSwipeMax`               | `BuyGoodsElasticSwipeMax`               |
| `AutoStockpileDevSkipSwipe`           | `BuyGoodsElasticDevSkipSwipe`           |
| `AutoStockpileRelayNodeSwipe`         | `BuyGoodsElasticRelayNodeSwipe`         |
| `AutoStockpileDevSkipBuy`             | `BuyGoodsElasticDevSkipBuy`             |
| `AutoStockpileBuy`                    | `BuyGoodsElasticBuy`                    |
| `AutoStockpileCheckBuyResult`         | `BuyGoodsElasticCheckBuyResult`         |
| `AutoStockpileCancelBuy`              | `BuyGoodsElasticCancelBuy`              |

### 2.6 `assets/resource/pipeline/AutoStockpile/Staple/Item.json`

| 当前节点名                           | 目标节点名                           |
| ------------------------------------ | ------------------------------------ |
| `AutoStockpileInStapleItem`          | `BuyGoodsStapleInItem`               |
| `AutoStockpileInStapleItemDiscounts` | `BuyGoodsStapleInItemDiscounts`      |
| `AutoStockpileInStapleItemName`      | `BuyGoodsStapleInItemName`           |
| `AutoStockpileTargetStockBillColor`  | `BuyGoodsStapleTargetStockBillColor` |
| `AutoStockpileTargetStockBillText`   | `BuyGoodsStapleTargetStockBillText`  |
| `AutoStockpileTargetStockBill`       | `BuyGoodsStapleTargetStockBill`      |
| `AutoStockpileTargetCanBuy`          | `BuyGoodsStapleTargetCanBuy`         |
| `AutoStockpileSwipeToMax`            | `BuyGoodsStapleSwipeToMax`           |
| `AutoStockpileSwipeBuy`              | `BuyGoodsStapleSwipeBuy`             |
| `AutoStockpileCurrentStockBill`      | `BuyGoodsStapleCurrentStockBill`     |

### 2.7 `assets/resource/pipeline/AutoStockpile/Staple/Wuling.json`

| 当前节点名                           | 目标节点名                            |
| ------------------------------------ | ------------------------------------- |
| `AutoStockpileInStapleWuling`        | `BuyGoodsStapleInWuling`              |
| `AutoStockpileTargetCanNotBuyWuling` | `BuyGoodsStapleTargetCanNotBuyWuling` |
| `AutoStockpileBuyItemWulingTask`     | `BuyGoodsStapleBuyItemWulingTask`     |
| `AutoStockpileSwipeWuling`           | `BuyGoodsStapleSwipeWuling`           |

### 2.8 `assets/resource/pipeline/AutoStockpile/Staple/ValleyIV.json`

| 当前节点名                             | 目标节点名                              |
| -------------------------------------- | --------------------------------------- |
| `AutoStockpileInStapleValleyIV`        | `BuyGoodsStapleInValleyIV`              |
| `AutoStockpileTargetCanNotBuyValleyIV` | `BuyGoodsStapleTargetCanNotBuyValleyIV` |
| `AutoStockpileBuyItemValleyIVTask`     | `BuyGoodsStapleBuyItemValleyIVTask`     |
| `AutoStockpileSwipeValleyIV`           | `BuyGoodsStapleSwipeValleyIV`           |

---

## 三、必须同步更新的外围引用

| 文件/范围                              | 当前引用                   | 目标引用                     | 说明                               |
| -------------------------------------- | -------------------------- | ---------------------------- | ---------------------------------- |
| `assets/tasks/AutoStockpile.json`      | 任务文件名/任务名          | `BuyGoods.json` / `BuyGoods` | 任务入口同步更名                   |
| `assets/interface.json`                | `tasks/AutoStockpile.json` | `tasks/BuyGoods.json`        | 接口入口同步                       |
| `assets/resource/image/AutoStockpile/` | `AutoStockpile/...`        | `BuyGoods/...`               | 仅更名图片文件夹，不更名图片文件名 |
| pipeline 中所有 `template` 引用        | `AutoStockpile/*.png`      | `BuyGoods/*.png`             | 只更新路径前缀，文件名保持不变     |
| `next` / `on_error` / `subtask`        | 所有旧节点名               | 对应新节点名                 | 所有跨节点跳转必须同步             |
| `custom_action` / `custom_recognition` | `AutoStockpile.*`          | `BuyGoodsElastic.*`          | Go 与 pipeline 契约需一致          |
| Go 常量中的节点名                      | `AutoStockpile*`           | `BuyGoodsElastic*`           | 仅限 Elastic 相关 Go 调用链        |

### 3.1 Image 同步规则

- `assets/resource/image/AutoStockpile/` 整体更名为 `assets/resource/image/BuyGoods/`。
- **仅更名文件夹名，不更名其中的图片文件名**。
- 所有 pipeline JSON 中的 `template` 路径同步从 `AutoStockpile/...` 更新为 `BuyGoods/...`。
- 若 Go 侧存在任何基于图片目录前缀拼接的路径，也需要同步从 `AutoStockpile` 更新为 `BuyGoods`，但图片 basename 保持不变。

---

## 四、实施顺序建议

1. 先改 Go 包目录与 `register.go` 引用。
2. 再改 Go 中的 custom action / recognition 名称与节点常量。
3. 修改 `Main.json` 中的总入口节点名。
4. 修改 `Elastic/*` 全部节点名与内部引用。
5. 修改 `Staple/*` 全部节点名与内部引用。
6. 修改 `assets/tasks/AutoStockpile.json` 与 `assets/interface.json`。
7. 同步更新 image 文件夹路径引用：`AutoStockpile/...` → `BuyGoods/...`，但不修改图片文件名。
8. 最后统一检查所有 `next` / `on_error` / `subtask` / `custom_*` / override / `template` 引用。

---

## 五、审计结论

- **Go 与 pipeline 前缀策略是刻意分离的**：
    - Go：`BuyGoodsElastic.*`
    - Pipeline Main：`BuyGoods*`
    - Pipeline Elastic：`BuyGoodsElastic*`
    - Pipeline Staple：`BuyGoodsStaple*`
- **本计划是完整更名计划，不含代码修改**。
- 真正实施时，重点风险在于：
    - Go 注册字符串与 pipeline `Custom` 节点不一致
    - `next` / `on_error` / `subtask` 的旧节点名遗漏
    - `assets/tasks/AutoStockpile.json` 中的 override 引用未同步
    - image 目录已更名但 pipeline `template` 仍指向旧的 `AutoStockpile/...` 路径
