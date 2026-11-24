# Meeting Rooms II

> NeetCode 150 · Intervals · Blind 75 — decouple starts from ends, sweep both sorted streams; peak concurrency is the room count. Go, stdlib only, tested.

## Problem

The fewest rooms that accommodate all meetings. Same mathematics as Minimum Railway Platforms (dsa-044) with one semantic flip: here a meeting ending at time t *does* free its room for a meeting starting at t — half-open intervals.

## Approach

Which meeting occupies which room never matters, only how many run at once: sort the starts and ends as independent streams and merge-sweep, incrementing on a start, decrementing on an end, tracking the peak. The boundary convention lives in a single character — `starts[i] < ends[j]` counts a start strictly before an end as concurrent, so equal timestamps process the end first and back-to-back meetings share a room (the platforms problem used `<=` for the opposite ruling; both repos exist so the pair documents the difference). The min-heap-of-end-times formulation is equivalent; the two-pointer sweep allocates less and reads cleaner.

**Complexity:** O(n log n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
