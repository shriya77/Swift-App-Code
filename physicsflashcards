// Adapted from: https://www.hackingwithswift.com/books/ios-swiftui/showing-and-hiding-views-with-transitions
// Student work
import SwiftUI

struct ContentView: View {
    @State private var anim_choice = 0
    let image_choose = ["ss1","ss2","ss3"]
    var body: some View {
        VStack {
            Button("Shuffle flashcards") {
                anim_choice = Int.random(in: 0...(image_choose.count - 1)) 
            }
            Image(image_choose[anim_choice])
                .resizable()
                .scaledToFit()
                .frame(width: 400, height: 300)
            //.clipShape(RoundedRectangle(cornerRadius: 10))
        }
    }
}
