
//({ key set}, value) 로 map
function m() {
   for(var i in this.detail) {
        emit({acn_model:this.acn_model,
              acn_code:this.detail[i].acn_code,
              credit_debit:this.detail[i].credit_debit,
            }, this.detail[i].detail_amount);
   }
}

//val에대한 reduce를 sum으로 수행
function r(key, val) {
     return Array.sum(val)
}

//특정회계전표일자에대한 MR수행하여 daily_closing collection에 저장
db.journals.mapReduce(m, r, {query: {acn_date:"20180128"}, out:"daily_closing"})


//결산결과 확인
db.daily_closing.find()
