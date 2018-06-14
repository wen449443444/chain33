package accounts

import (

	"time"

	l "github.com/inconshreveable/log15"

	"gitlab.33.cn/chain33/chain33/types"

	//"encoding/json"
	//"io/ioutil"
	"strconv"
	"fmt"
	"gitlab.33.cn/chain33/chain33/rpc"
	"os"
)

const secondsPerBlock = 15
const btyPreBlock = 18

var log = l.New("module", "accounts")

type ShowMinerAccount struct {
	DataDir string
	Addrs []string
}

func (*ShowMinerAccount) Echo(in *string, out *interface{}) error {
	if in == nil {
		return types.ErrInputPara
	}
	*out = *in
	return nil
}

type TimeAt struct {
	// YYYY-mm-dd-HH
	TimeAt string `json:"timeAt"`
}

func (show *ShowMinerAccount) Get(in *TimeAt, out *interface{}) error {
	if in == nil {
		log.Error("show", "in", "nil")
		return types.ErrInputPara
	}
	seconds := time.Now().Unix()
	if len(in.TimeAt) != 0 {
		tm, err := time.Parse("2006-01-02-15", in.TimeAt)
		if err != nil {
			log.Error("show", "in.TimeAt Parse", err)
			return types.ErrInputPara
		}
		seconds = tm.Unix()
	}
	log.Info("show", "utc", seconds)

	addrs := show.Addrs
	log.Info("show", "miners", addrs)
	//for i := int64(0); i < 200; i++ {
		header, curAcc, err := cache.getBalance(addrs, "ticket", seconds)
		if err != nil {
			return nil
		}
		lastHourHeader, lastAcc, err := cache.getBalance(addrs, "ticket", header.BlockTime-3600)
		if err != nil {
			return nil
		}
		fmt.Print(curAcc, lastAcc)

		miner := &MinerAccounts{}
		miner.Seconds = header.BlockTime - lastHourHeader.BlockTime
		miner.Blocks = header.Height - lastHourHeader.Height
		miner.ExpectBlocks = miner.Seconds / secondsPerBlock

		miner = calcResult(miner, curAcc, lastAcc)
		*out = &miner

		seconds = seconds - 3600
	//}


	return nil
}

func calcResult(miner *MinerAccounts, acc1, acc2 []*rpc.Account) *MinerAccounts {
	type minerAt struct {
		addr string
		curAcc *rpc.Account
		lastAcc *rpc.Account
	}
	miners := map[string]*minerAt{}
	for _, a := range acc1 {
		miners[a.Addr] = &minerAt{a.Addr, a, nil}
	}
	for _, a := range acc2 {
		if _, ok := miners[a.Addr]; ok {
			miners[a.Addr].lastAcc = a
		}
	}

	totalIncrease := float64(0)
	totalFrozen := float64(0)
	for _, v := range miners {
		if v.lastAcc != nil && v.curAcc != nil {
			totalFrozen += float64(v.curAcc.Frozen)/float64(types.Coin)
		}
	}
	for k, v := range miners {
		if v.lastAcc != nil && v.curAcc != nil {
			total := v.curAcc.Balance + v.curAcc.Frozen
			increase := total - v.lastAcc.Balance - v.lastAcc.Frozen
			expectIncrease := float64(miner.ExpectBlocks) * float64(btyPreBlock) * (float64(v.curAcc.Frozen)/float64(types.Coin)) / totalFrozen

			m :=  &MinerAccount{
				Addr: k,
				Total: strconv.FormatFloat(float64(total)/float64(types.Coin), 'f', 4, 64),
				Increase: strconv.FormatFloat(float64(increase)/float64(types.Coin), 'f', 4, 64),
				Frozen: strconv.FormatFloat(float64(v.curAcc.Frozen)/float64(types.Coin), 'f', 4, 64),
				ExpectIncrease: strconv.FormatFloat(expectIncrease, 'f', 4, 64),
			}

			//if m.Addr == "1Lw6QLShKVbKM6QvMaCQwTh5Uhmy4644CG" {
			//	log.Info("acount", "Increase", m.Increase, "expectIncrease", m.ExpectIncrease)
			//	fmt.Println(os.Stderr, "Increase", m.Increase, "expectIncrease", m.ExpectIncrease)
			//}

			miner.MinerAccounts = append(miner.MinerAccounts, m)
			totalIncrease += float64(increase)/float64(types.Coin)
		}
	}
	miner.TotalIncrease = strconv.FormatFloat(totalIncrease, 'f', 4, 64)

	return miner

}

/*
func readJson(jsonFile string) (*Accounts, error) {
	d1, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Error("show", "read json", jsonFile, "err", err)
		return nil, err
	}
	var acc Accounts
	err = json.Unmarshal([]byte(d1), &acc)
	if err != nil {
		log.Error("show", "read json", jsonFile, "err", err)
		return nil, err
	}
	return &acc, nil
}

func parseAccounts(acc *Accounts) (*map[string]float64, error) {
	result := map[string]float64{}
	for _, a := range acc.Accounts {
		f1, e1 := strconv.ParseFloat(a.Balance, 64)
		f2, e2 := strconv.ParseFloat(a.Frozen, 64)
		if e1 != nil || e2 != nil {
			log.Error("show", "account2  len", e1, "account  len", e2)
			return nil, types.ErrNotFound
		}
		result[a.Addr] = f1 + f2
	}
	return &result, nil
}

func getAccountDetail(jsonFile string) (*map[string]float64, error) {
	acc, err := readJson(jsonFile)
	if err != nil {
		return nil, err
	}
	return parseAccounts(acc)
}

*/

