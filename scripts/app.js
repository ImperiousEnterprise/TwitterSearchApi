var vm = new Vue({
    el: '#vue-instance',
    delimiters: ['${', '}'],
    data:{
        text: '',
        tweets: [],
        errout: '',
    },
    filters:{
        highlight: function(words, query){

            if(words.indexOf("#") > -1){
                words = words.replace( /(^|\s)#([a-z\d]+)/ig, "$1<a href=\'https://twitter.com/hashtag/$2?src=hash'>#$2</a>");
            }

            if(words.indexOf("@") > -1){
                words = words.replace( /(^|\s)@([a-zA-Z\d]+)/ig, "$1<a href=\'https://twitter.com/$2'>@$2</a>");
            }

            query.split(/[\\\s_+-.,!@#$%^&*();\/|<>"']+/g).forEach(function(el){
                if(el.length < 1) return;
                words = words.replace(new RegExp("(?!\/)("+el+")(?![^<]*>)", "g"), '<b>' + el + '</b>');
            });
            return words;
        }
    },
    watch: {
        text: function (val) {
            if(val.length < 1){
                this.tweets = [];
            }
        }
    },
    methods:{
        SearchTweets: function (e){
            this.$http.get('/search?q=' + encodeURIComponent(this.text) +'&count=20')
                .then( response => { this.tweets = response.data; this.errout=''; this.SortByCreatedAt;})
                .catch(err => {this.tweets = []; this.errout = err.body;});
        },
        SortByCreatedAt: function (){
                this.tweets.sort(function(left, right){
                    return moment(right.created_at).diff(moment(left.created_at)); // No more need to convert strings to dates
            })
        },
        GenerateUserUrl: function (e){
            return "https://twitter.com/"+e;
        },
        SwitchToTimeAgo: function (time) {
            if(!moment(time, 'ddd MMM DD HH:mm:ss ZZ YYYY', 'en').isValid()){
                var MonthToNumber = {
                  //You need to create convert months to numbers
                };
               var split = time.split(" ");
               var spliTime = split[3].split(":");
               time = new Date(split[5],MonthToNumber[split[1]],split[2],spliTime[0],spliTime[1],spliTime[2]);

            }
            return moment(time, 'ddd MMM DD HH:mm:ss ZZ YYYY', 'en').fromNow();
        }
    }
});


