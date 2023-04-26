package texdocs.slides.sections.bigo.code;

import java.util.List;

public class ListSearch {
    /**
     * Search for the largest value in the list.
     */
    public static int searchMax(List<Integer> list) {
        int max = Integer.MIN_VALUE;
        for (int i = 0; i < list.size(); i++) {
            max = Math.max(max, list.get(i));
        }
        return max;
    }

    /** 
     * Search for the smallest and largest value in the list and return their difference.
     */
    public static int diffMinMax(List<Integer> list) {
        int min = Integer.MAX_VALUE;
        int max = Integer.MIN_VALUE;
        for (int i = 0; i < list.size(); i++) {
            min = Math.min(min, list.get(i));
        }
        for (int i = 0; i < list.size(); i++) {
            max = Math.max(max, list.get(i));
        }
        return max - min;
    }

    /** 
     * Search for the pair of values in the list that have the minimal difference.
     */
    public static int closestPair(List<Integer> list) {
        int minDiff = Integer.MAX_VALUE;
        for (int i = 0; i < list.size(); i++) {
            for (int j = i + 1; j < list.size(); j++) {
                minDiff = Math.min(minDiff, Math.abs(list.get(i) - list.get(j)));
            }
        }
        return minDiff;
    }
}
