// Adapted from: https://www.hackingwithswift.com/books/ios-swiftui/showing-and-hiding-views-with-transitions
// Student work

import SwiftUI

struct ContentView: View {
    @State private var anim_choice = 0
    let image_choose = ["card1","card2","card3","card4","card5"]
    var body: some View {
        VStack {
            Button("GENERATE ANOTHER ONE") {
                anim_choice = Int.random(in: 0...(image_choose.count - 1)) 
            }
            Image(image_choose[anim_choice])
                .resizable()
                .scaledToFit()
                .frame(width: 450, height: 500)
            //.clipShape(RoundedRectangle(cornerRadius: 10))
        }
    }
}
