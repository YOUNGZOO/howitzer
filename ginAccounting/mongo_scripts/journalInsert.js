//mongo client 에서 sample data put 용도

//drop all
db.journals.drop();

//예시전표1 insert
db.journals.insert(
{
  acn_date:"20180128",
  acn_model:"00001",
  acn_amount:12345.67,
  detail:[
  	{
    	acn_code:"1000000001",
    	credit_debit:true,
    	detail_amount:12345.67
  	},
  	{
    	acn_code:"2000000001",
    	credit_debit:false,
    	detail_amount:12000
  	},
  	{
    	acn_code:"2000000002",
    	credit_debit:false,
    	detail_amount:345.67
  	}
  ],
  timestamp: Date()
});

//예시전표2 insert
db.journals.insert(
{
  acn_date:"20180128",
  acn_model:"00002",
  acn_amount:12000,
  detail:[
  	{
    	acn_code:"2000000001",
    	credit_debit:true,
    	detail_amount:12000
  	},
  	{
    	acn_code:"3000000001",
    	credit_debit:false,
    	detail_amount:10000
  	},
  	{
    	acn_code:"3000000002",
    	credit_debit:false,
    	detail_amount:2000
  	}
  ],
  timestamp: Date()
});

//예시전표 조회
db.journals.find();
