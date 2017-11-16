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
    methods:{
        SearchTweets: function (e){
            this.$http.get('/search?q=' + encodeURIComponent(this.text) +'&count=20')
                .then( response => { this.tweets = response.data; this.errout=''; })
                .catch(err => {this.tweets = []; this.errout = err.body;});
        },
        GenerateUserUrl: function (e){
            return "https://twitter.com/"+e;
        },
        SwitchToTimeAgo: function (time) {
            if(!moment(time, 'ddd MMM DD HH:mm:ss ZZ YYYY', 'en').isValid()){
                var MonthToNumber = {
                    Jan: 0,
                    Feb: 1,
                    Mar: 2,
                    Apr: 3,
                    May: 4,
                    Jun: 5,
                    Jul: 6,
                    Aug: 7,
                    Sep: 8,
                    Oct: 9,
                    Nov: 10,
                    Dec: 11
                };
               var split = time.split(" ");
               var spliTime = split[3].split(":");
               time = new Date(split[5],MonthToNumber[split[1]],split[2],spliTime[0],spliTime[1],spliTime[2]);

            }
            return moment(time, 'ddd MMM DD HH:mm:ss ZZ YYYY', 'en').fromNow();
        }
    }
});


