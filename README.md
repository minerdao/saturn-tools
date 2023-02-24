# saturn-tools
#### 这是一个saturn-node的运维工具,提供的功能,适合有多个地址用户使用,主要提供以下功能:
- 统计节点最近24小时收益
- 统计最近24小时所有地址总收益
#### 所有数据都来源于[https://dashboard.strn.network/stats],
### 1、使用此工具前需要安装go:
```
apt  install gccgo-go
```
### 2、下载代码:
```
git clone https://github.com/minerdao/saturn-tools.git
```
### 3、使用代码前请修改代码中地址为需要统计的地址,涉及到三处修改:
```
var filAddressLocationMap = map[string]string{
        "f1wf7lu7quwz5hgsl5qybnjf................": "US", //第二个字段代表地址对应的地理位置,根据自己服务器地址实际位置填写
        "f1pyfpcscyfto7phqkv2ybx7................": "US",
        "f1qbqd757pul7b5dpttmcz2c................": "US",
========================================================================
func main() {
        filAddresses := []string{
                "f1wf7lu7quwz5hgsl........................", //这里对应的是上文中的地址
                "f1pyfpcscyfto7phq........................",
                "f1qbqd757pul7b5dp........................",
        }
        
=============================================================================
func sumTotalFilAmount() {
        filAddresses := []string{
                "f1wf7lu7quwz5hgsl5qyb................", //这里对应的是上文中的地址
                "f1pyfpcscyfto7phqkv2y................",
                "f1qbqd757pul7b5dpttmc................",
        }



```
### 4、运行程序:
```
go run check_reward.go
```
