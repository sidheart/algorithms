/*
You are given two string arrays creators and ids, and an integer array views, all of length n. The ith video on a platform was created by creator[i], has an id of ids[i], and has views[i] views.

The popularity of a creator is the sum of the number of views on all of the creator's videos. Find the creator with the highest popularity and the id of their most viewed video.

    If multiple creators have the highest popularity, find all of them.
    If multiple videos have the highest view count for a creator, find the lexicographically smallest id.

Return a 2D array of strings answer where answer[i] = [creatori, idi] means that creatori has the highest popularity and idi is the id of their most popular video. The answer can be returned in any order.
*/
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;

class Solution {
    private static class Video {
        public String id;
        public int views;

        public Video(String id, int views) {
            this.id = id;
            this.views = views;
        }
    }

    private static class Creator {
        public String id;
        public int totalViews;
        public Video bestVideo;

        public Creator(String id) {
            this.id = id;
            this.totalViews = 0;
            this.bestVideo = null;
        }

        public void addVideo(Video v) {
            totalViews += v.views;
            if (bestVideo == null || v.views > bestVideo.views || (bestVideo.views == v.views && v.id.compareTo(bestVideo.id) < 0)) {
                bestVideo = v;
            }
        }
    }

    public List<List<String>> mostPopularCreator(String[] creators, String[] ids, int[] views) {
        // process creators and videos
        Map<String, Creator> creatorMap = new HashMap<>();
        for (int i = 0; i < creators.length; i++) {
            String creatorId = creators[i];
            Creator c;
            if (creatorMap.containsKey(creatorId)) {
                c = creatorMap.get(creatorId);
            } else {
                c = new Creator(creatorId);
                creatorMap.put(creatorId, c);
            }
            c.addVideo(new Video(ids[i], views[i]));
        }
        // find the best creators
        List<Creator> topCreators = new ArrayList<>();
        creatorMap.forEach( (id, c) -> {
            if (topCreators.isEmpty() || topCreators.get(0).totalViews == c.totalViews) {
                topCreators.add(c);
            } else if (c.totalViews > topCreators.get(0).totalViews) {
                topCreators.clear();
                topCreators.add(c);
            }
        });
        // format the output
        List<List<String>> answer = new ArrayList<>();
        for (Creator c : topCreators) {
            answer.add(Arrays.asList(c.id, c.bestVideo.id));
        }
        return answer;
    }
}
